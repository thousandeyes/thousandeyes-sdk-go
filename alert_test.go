package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_GetAlertRule(t *testing.T) {
	out := `{"ruleId":1, "ruleName": "test", "roundsViolatingOutOf": 1, "roundsViolatingRequired": 1}`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/alert-rules/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := AlertRule{RuleId: 1, RuleName: "test", RoundsViolatingOutOf: 1, RoundsViolatingRequired: 1}

	res, err := client.GetAlertRule(1)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteAlertRule(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/alert-rules/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteAlertRule(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateAlertRule(t *testing.T) {
	setup()
	out := `{"ruleId":1, "ruleName": "test", "roundsViolatingOutOf": 2, "roundsViolatingRequired": 1}`
	mux.HandleFunc("/alert-rules/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	u := AlertRule{RoundsViolatingOutOf: 2}
	res, err := client.UpdateAlertRule(id, u)
	if err != nil {
		t.Fatal(err)
	}
	expected := AlertRule{RuleId: 1, RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
	assert.Equal(t, &expected, res)
}

func TestClient_CreateAlertRule(t *testing.T) {
	setup()
	out := `{"ruleId":1, "ruleName": "test", "roundsViolatingOutOf": 2, "roundsViolatingRequired": 1}`
	mux.HandleFunc("/alert-rules/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	u := AlertRule{RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
	res, err := client.CreateAlertRule(u)
	if err != nil {
		t.Fatal(err)
	}
	expected := AlertRule{RuleId: 1, RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
	assert.Equal(t, &expected, res)
}
