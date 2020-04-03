package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetBGPMonitors(t *testing.T) {
	out := ` { "bgpMonitors": [ {"monitorId":1, "monitorType": "bgp","monitorName": "test", "ipAddress": "1.2.3.4"} ] }`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/bgp-monitors.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := BGPMonitors{
		BGPMonitor{MonitorID: 1, MonitorName: "test", MonitorType: "bgp", IPAddress: "1.2.3.4"},
	}

	res, err := client.GetBPGMonitors()
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetBGPMonitorsAlertError(t *testing.T) {

	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
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
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/bgp-monitors.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetBPGMonitors()
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}
