package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_GetTestStatusCode(t *testing.T) {
	setup()
	out := `{"tests":[]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetTest(1)
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_GetTestsStatusCode(t *testing.T) {
	setup()
	out := `{"tests":[]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetTests()
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_GetGenericTestsJsonError(t *testing.T) {
	out := `{"tests": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetTests()
	assert.Error(t, err)
	assert.EqualError(t, err, "could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GetGenericTestJsonError(t *testing.T) {
	out := `{"tests": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetTest(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}
