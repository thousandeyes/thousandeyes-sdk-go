package request

import (
	"errors"
	"net/url"
	"reflect"
	"testing"
	"time"
)

type mappedValue struct {
	values map[string]interface{}
	err    error
}

func (m mappedValue) ToMap() (map[string]interface{}, error) {
	return m.values, m.err
}

func TestSelectHeaders(t *testing.T) {
	if got := SelectHeaderContentType(nil); got != "" {
		t.Fatalf("SelectHeaderContentType(nil) = %q, want empty", got)
	}
	if got := SelectHeaderContentType([]string{"text/plain", "application/json"}); got != "application/json" {
		t.Fatalf("SelectHeaderContentType() = %q, want application/json", got)
	}
	if got := SelectHeaderContentType([]string{"text/plain", "application/xml"}); got != "text/plain" {
		t.Fatalf("SelectHeaderContentType() = %q, want first value", got)
	}
	if got := SelectHeaderAccept([]string{"application/json", "text/plain"}); got != "application/json,text/plain" {
		t.Fatalf("SelectHeaderAccept() = %q, want joined values", got)
	}
}

func TestParameterValueToString(t *testing.T) {
	if got := ParameterValueToString(42, "value"); got != "42" {
		t.Fatalf("ParameterValueToString(scalar) = %q, want %q", got, "42")
	}

	mapped := &mappedValue{values: map[string]interface{}{"value": "mapped"}}
	if got := ParameterValueToString(mapped, "value"); got != "mapped" {
		t.Fatalf("ParameterValueToString(mapped) = %q, want %q", got, "mapped")
	}

	failing := &mappedValue{err: errors.New("mapping failed")}
	if got := ParameterValueToString(failing, "value"); got != "" {
		t.Fatalf("ParameterValueToString(failing) = %q, want empty", got)
	}
}

func TestParameterAddToQuery(t *testing.T) {
	values := url.Values{}
	timestamp := time.Date(2026, time.July, 24, 12, 30, 45, 123, time.UTC)

	ParameterAddToHeaderOrQuery(values, "string", "value", "")
	ParameterAddToHeaderOrQuery(values, "integer", int64(-42), "")
	ParameterAddToHeaderOrQuery(values, "unsigned", uint(42), "")
	ParameterAddToHeaderOrQuery(values, "float", float64(1.5), "")
	ParameterAddToHeaderOrQuery(values, "boolean", true, "")
	ParameterAddToHeaderOrQuery(values, "repeated", []string{"one", "two"}, "")
	ParameterAddToHeaderOrQuery(values, "csv", []string{"one", "two"}, "csv")
	ParameterAddToHeaderOrQuery(values, "filter", map[string]interface{}{"name": "example"}, "")
	ParameterAddToHeaderOrQuery(values, "timestamp", timestamp, "")
	ParameterAddToHeaderOrQuery(values, "nil", nil, "")

	want := url.Values{
		"string":       {"value"},
		"integer":      {"-42"},
		"unsigned":     {"42"},
		"float":        {"1.5"},
		"boolean":      {"true"},
		"repeated":     {"one", "two"},
		"csv":          {"one,two"},
		"filter[name]": {"example"},
		"timestamp":    {timestamp.Format(time.RFC3339Nano)},
		"nil":          {"null"},
	}
	if !reflect.DeepEqual(values, want) {
		t.Fatalf("serialized query = %#v, want %#v", values, want)
	}
}

func TestParameterAddToHeader(t *testing.T) {
	headers := map[string]string{}
	ParameterAddToHeaderOrQuery(headers, "X-Count", 3, "")

	if got := headers["X-Count"]; got != "3" {
		t.Fatalf("X-Count = %q, want %q", got, "3")
	}
}

func TestParameterToJSON(t *testing.T) {
	got, err := ParameterToJson(map[string]string{"name": "example"})
	if err != nil {
		t.Fatalf("ParameterToJson() error = %v", err)
	}
	if got != `{"name":"example"}` {
		t.Fatalf("ParameterToJson() = %q, want JSON object", got)
	}

	if _, err := ParameterToJson(make(chan int)); err == nil {
		t.Fatal("ParameterToJson(channel) error = nil, want error")
	}
}
