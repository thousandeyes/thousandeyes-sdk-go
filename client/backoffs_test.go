package client

import (
	"net/http"
	"testing"
	"time"
)

func useFixedTime(t *testing.T, now time.Time) {
	t.Helper()

	original := timeNow
	timeNow = func() time.Time { return now }
	t.Cleanup(func() {
		timeNow = original
	})
}

func TestParseRetryAfterHeader(t *testing.T) {
	now := time.Date(2026, time.July, 24, 12, 0, 0, 0, time.UTC)
	useFixedTime(t, now)

	tests := []struct {
		name    string
		headers []string
		want    time.Duration
		wantOK  bool
	}{
		{name: "missing", wantOK: false},
		{name: "empty", headers: []string{""}, wantOK: false},
		{name: "seconds", headers: []string{"120"}, want: 120 * time.Second, wantOK: true},
		{name: "zero seconds", headers: []string{"0"}, wantOK: true},
		{name: "negative seconds", headers: []string{"-1"}, wantOK: false},
		{name: "overflow seconds", headers: []string{"9999999999"}, wantOK: false},
		{name: "invalid", headers: []string{"later"}, wantOK: false},
		{
			name:    "IMF-fixdate",
			headers: []string{now.Add(2 * time.Minute).Format(http.TimeFormat)},
			want:    2 * time.Minute,
			wantOK:  true,
		},
		{
			name:    "RFC850",
			headers: []string{now.Add(2 * time.Minute).Format(time.RFC850)},
			want:    2 * time.Minute,
			wantOK:  true,
		},
		{
			name:    "ANSI C asctime",
			headers: []string{now.Add(2 * time.Minute).Format(time.ANSIC)},
			want:    2 * time.Minute,
			wantOK:  true,
		},
		{
			name:    "past date",
			headers: []string{now.Add(-time.Minute).Format(http.TimeFormat)},
			wantOK:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := parseRetryAfterHeader(test.headers)
			if ok != test.wantOK {
				t.Fatalf("parseRetryAfterHeader() ok = %v, want %v", ok, test.wantOK)
			}
			if got != test.want {
				t.Fatalf("parseRetryAfterHeader() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestParseResetHeader(t *testing.T) {
	now := time.Unix(1_800_000_000, 0)
	useFixedTime(t, now)

	future := "1800000120"
	past := "1799999940"
	whitespace := " 1800000120 "
	invalid := "not-a-timestamp"
	overflow := "999999999999999999999999"

	tests := []struct {
		name   string
		value  *string
		want   time.Duration
		wantOK bool
	}{
		{name: "nil", wantOK: false},
		{name: "future", value: &future, want: 2 * time.Minute, wantOK: true},
		{name: "past", value: &past, wantOK: true},
		{name: "surrounding whitespace", value: &whitespace, wantOK: false},
		{name: "invalid", value: &invalid, wantOK: false},
		{name: "overflow", value: &overflow, wantOK: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := parseResetHeader(test.value)
			if ok != test.wantOK {
				t.Fatalf("parseResetHeader() ok = %v, want %v", ok, test.wantOK)
			}
			if got != test.want {
				t.Fatalf("parseResetHeader() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestThousandEyesBackoff(t *testing.T) {
	now := time.Unix(1_800_000_000, 0)
	useFixedTime(t, now)

	const (
		min     = time.Second
		max     = 30 * time.Second
		attempt = 2
	)
	futureOrganization := "1800000120"
	futureInstant := "1800000060"
	past := "1799999940"

	response := func(status int, headers map[string]string) *http.Response {
		header := make(http.Header)
		for name, value := range headers {
			header.Set(name, value)
		}
		return &http.Response{StatusCode: status, Header: header}
	}

	tests := []struct {
		name string
		resp *http.Response
		want time.Duration
	}{
		{name: "nil response", want: 0},
		{
			name: "Retry-After seconds on 429",
			resp: response(http.StatusTooManyRequests, map[string]string{"Retry-After": "10"}),
			want: 10 * time.Second,
		},
		{
			name: "Retry-After seconds on 503",
			resp: response(http.StatusServiceUnavailable, map[string]string{"Retry-After": "10"}),
			want: 10 * time.Second,
		},
		{
			name: "Retry-After HTTP-date",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"Retry-After": now.Add(90 * time.Second).Format(http.TimeFormat),
			}),
			want: 90 * time.Second,
		},
		{
			name: "Retry-After ignored for other status",
			resp: response(http.StatusOK, map[string]string{"Retry-After": "10"}),
			want: 4 * time.Second,
		},
		{
			name: "Retry-After precedes reset headers",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"Retry-After":                     "10",
				"x-organization-rate-limit-reset": futureOrganization,
			}),
			want: 10 * time.Second,
		},
		{
			name: "organization reset precedes instant-test reset",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"x-organization-rate-limit-reset": futureOrganization,
				"x-instant-test-rate-limit-reset": futureInstant,
			}),
			want: 2 * time.Minute,
		},
		{
			name: "invalid organization reset falls through to instant-test reset",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"x-organization-rate-limit-reset": "invalid",
				"x-instant-test-rate-limit-reset": futureInstant,
			}),
			want: time.Minute,
		},
		{
			name: "invalid Retry-After falls through to reset header",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"Retry-After":                     "invalid",
				"x-organization-rate-limit-reset": futureOrganization,
			}),
			want: 2 * time.Minute,
		},
		{
			name: "overflowing Retry-After falls through to reset header",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"Retry-After":                     "9999999999",
				"x-organization-rate-limit-reset": futureOrganization,
			}),
			want: 2 * time.Minute,
		},
		{
			name: "overflowing Retry-After without reset falls back",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"Retry-After": "9999999999",
			}),
			want: 4 * time.Second,
		},
		{
			name: "past reset returns zero",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"x-organization-rate-limit-reset": past,
			}),
			want: 0,
		},
		{
			name: "invalid reset falls back",
			resp: response(http.StatusTooManyRequests, map[string]string{
				"x-organization-rate-limit-reset": "invalid",
			}),
			want: 4 * time.Second,
		},
		{
			name: "missing headers falls back",
			resp: response(http.StatusTooManyRequests, nil),
			want: 4 * time.Second,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := thousandEyesBackoff(min, max, attempt, test.resp); got != test.want {
				t.Fatalf("thousandEyesBackoff() = %v, want %v", got, test.want)
			}
		})
	}
}
