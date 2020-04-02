package thousandeyes

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetBGPMonitors(t *testing.T) {
	out := ` { "bgpMonitors": [ {"monitorId":1, "monitorType": "bgp","monitorName": "test", "ipAddress": "1.2.3.4"} ] }`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/bgp-monitors.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := BGPMonitors{
		BGPMonitor{MonitorID: 1, MonitorName: "test", MonitorType: "bgp", IPAddress: "1.2.3.4"},
	}

	res, err := client.GetBPGMonitors()
	fmt.Println(res)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetBGPMonitorsAlertError(t *testing.T) {

	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/bgp-monitors.json", func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetBPGMonitors()
	teardown()
	assert.Error(t, err)

}

func TestClient_GetBPGMonitorsJsonError(t *testing.T) {
	out := ` { "bgpMonitors": [ {aonitorId":1, "monitorType": "bgp","monitorName": "test", "ipAddress": "1.2.3.4"} ] }`

	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/bgp-monitors.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetBPGMonitors()
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}

// func TestClient_GetBGPMonitorsAlertError(t *testing.T) {
// 	setup()
// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	mux.HandleFunc("/alert-rules/1.json", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "GET", r.Method)
// 		w.WriteHeader(http.StatusBadRequest)
// 	})

// 	_, err := client.GetBGPMonitor(1)
// 	teardown()
// 	assert.Error(t, err)
// }

// func TestClient_DeleteBGPMonitor(t *testing.T) {
// 	setup()
// 	defer teardown()
// 	mux.HandleFunc("/alert-rules/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusNoContent)
// 		assert.Equal(t, "POST", r.Method)
// 	})

// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	id := 1
// 	err := client.DeleteBGPMonitor(id)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestClient_UpdateBGPMonitor(t *testing.T) {
// 	setup()
// 	out := `{"ruleId":1, "ruleName": "test", "roundsViolatingOutOf": 2, "roundsViolatingRequired": 1}`
// 	mux.HandleFunc("/alert-rules/1/update.json", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "POST", r.Method)
// 		_, _ = w.Write([]byte(out))
// 	})

// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	id := 1
// 	u := BGPMonitor{RoundsViolatingOutOf: 2}
// 	res, err := client.UpdateBGPMonitor(id, u)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	expected := BGPMonitor{RuleId: 1, RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
// 	assert.Equal(t, &expected, res)
// }

// func TestClient_CreateBGPMonitor(t *testing.T) {
// 	setup()
// 	out := `{"ruleId":1, "ruleName": "test", "roundsViolatingOutOf": 2, "roundsViolatingRequired": 1}`
// 	mux.HandleFunc("/alert-rules/new.json", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "POST", r.Method)
// 		w.WriteHeader(http.StatusCreated)
// 		_, _ = w.Write([]byte(out))
// 	})

// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	u := BGPMonitor{RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
// 	res, err := client.CreateBGPMonitor(u)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	expected := BGPMonitor{RuleId: 1, RuleName: "test", RoundsViolatingOutOf: 2, RoundsViolatingRequired: 1}
// 	assert.Equal(t, &expected, res)
// }

// func TestClient_AlertJsonError(t *testing.T) {
// 	out := `{"alertRules": [test]}`
// 	setup()
// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	mux.HandleFunc("/alert-rules/1.json", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "GET", r.Method)
// 		_, _ = w.Write([]byte(out))
// 	})
// 	_, err := client.GetBGPMonitor(1)
// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
// }

// func TestClient_GetAlertStatusCode(t *testing.T) {
// 	setup()
// 	out := `{"test":[{"testId":1,"testName":"test123"}]}`
// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	mux.HandleFunc("/alert-rules/1.json", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "GET", r.Method)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(out))
// 	})

// 	_, err := client.GetBGPMonitor(1)
// 	teardown()
// 	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
// }

// func TestClient_CreateAlertStatusCode(t *testing.T) {
// 	setup()
// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	mux.HandleFunc("/alert-rules/new.json", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "POST", r.Method)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(`{}`))
// 	})
// 	_, err := client.CreateBGPMonitor(BGPMonitor{})
// 	teardown()
// 	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
// }

// func TestClient_UpdateBGPMonitorStatusCode(t *testing.T) {
// 	setup()
// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	mux.HandleFunc("/alert-rules/1/update.json", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "POST", r.Method)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(`{}`))
// 	})
// 	_, err := client.UpdateBGPMonitor(1, BGPMonitor{})
// 	teardown()
// 	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
// }

// func TestClient_DeleteBGPMonitorStatusCode(t *testing.T) {
// 	setup()
// 	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
// 	mux.HandleFunc("/alert-rules/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, "POST", r.Method)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(`{}`))
// 	})
// 	err := client.DeleteBGPMonitor(1)
// 	teardown()
// 	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
// }
