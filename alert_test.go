package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAlertRule(t *testing.T) {
	out := `{"alertRules" : [ {"RuleID":1, "ruleName": "test" }]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/alert-rules.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := AlertRules{
		AlertRule{RuleID: 1, RuleName: "test"},
	}

	res, err := client.GetAlertRules()
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAlertError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/alert-rules.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetAlertRules()
	teardown()
	assert.Error(t, err)
}

func TestClient_DeleteAlertRule(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/alert-rules/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteAlertRule(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateAlertRule(t *testing.T) {
	setup()
	out := `{"RuleID":1, "ruleName": "test", "roundsViolatingOutOf": 2, "roundsViolatingRequired": 1}`
	mux.HandleFunc("/alert-rules/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	u := AlertRule{RoundsViolatingOutOf: 2}
	res, err := client.UpdateAlertRule(id, u)
	if err != nil {
		t.Fatal(err)
	}
	expected := AlertRule{RuleID: 1, RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
	assert.Equal(t, &expected, res)
}

func TestClient_CreateAlertRule(t *testing.T) {
	setup()
	out := `{"AlertRuleID":1, "ruleName": "test", "roundsViolatingOutOf": 2, "roundsViolatingRequired": 1}`
	mux.HandleFunc("/alert-rules/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	u := AlertRule{RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
	res, err := client.CreateAlertRule(u)
	if err != nil {
		t.Fatal(err)
	}
	expected := AlertRule{RuleID: 1, RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
	assert.Equal(t, &expected, res)
}

func TestClient_AlertJsonError(t *testing.T) {
	out := `{"alertRules": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/alert-rules.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAlertRules()
	assert.Error(t, err)
	assert.EqualError(t, err, "could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GetAlertStatusCode(t *testing.T) {
	setup()
	out := `{"alertRules":[{"ruleId":1,"ruleName":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/alert-rules.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetAlertRules()
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_CreateAlertStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/alert-rules/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreateAlertRule(AlertRule{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_UpdateAlertRuleStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/alert-rules/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdateAlertRule(1, AlertRule{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_DeleteAlertRuleStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/alert-rules/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeleteAlertRule(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}
