package client

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

var (
	rateLimitResetHeaders = []string{
		"x-organization-rate-limit-reset",
		"x-instant-test-rate-limit-reset",
	}
	resetHeaderPattern = regexp.MustCompile(`^\s*[0-9]+\s*$`)
	timeNow            = time.Now
)

// Custom backoff function
func customBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}

	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
		if sleep, ok := parseRetryAfterHeader(resp.Header["Retry-After"]); ok {
			return sleep
		}
	}

	for _, header := range rateLimitResetHeaders {
		if resetTimeStr := resp.Header.Get(header); resetTimeStr != "" {
			if sleep, ok := parseResetHeader(&resetTimeStr); ok {
				return sleep
			}

			// Default backoff if no valid reset time is found
			return retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
		}
	}

	// Default case for no Retry-After and rateLimitResetHeaders
	return retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
}

func parseRetryAfterHeader(headers []string) (time.Duration, bool) {
	if len(headers) == 0 || headers[0] == "" {
		return 0, false
	}
	header := headers[0]
	// Retry-After: 120
	if sleep, err := strconv.ParseInt(header, 10, 64); err == nil {
		if sleep < 0 { // a negative sleep doesn't make sense
			return 0, false
		}
		return time.Second * time.Duration(sleep), true
	}

	// Retry-After: Fri, 31 Dec 1999 23:59:59 GMT
	retryTime, err := time.Parse(time.RFC1123, header)
	if err != nil {
		return 0, false
	}
	if until := retryTime.Sub(timeNow()); until > 0 {
		return until, true
	}
	// date is in the past
	return 0, true
}

// Function to parse the reset header
func parseResetHeader(value *string) (time.Duration, bool) {
	if value == nil || !resetHeaderPattern.MatchString(*value) {
		return 0, false
	}

	// Parse the header value to a Unix timestamp
	resetTimeUnix, err := strconv.ParseInt(*value, 10, 64)
	if err != nil {
		return 0, false
	}

	// Calculate the duration until the reset time
	resetTime := time.Unix(resetTimeUnix, 0)
	waitDuration := time.Until(resetTime)

	// Return 0 if the reset time has already passed
	if waitDuration < 0 {
		return 0, true
	}
	return waitDuration, true
}
