package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_GetAgents(t *testing.T) {
	out := `{"agents":[{"agentId": 1}, {"agentId": 2}]}`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := Agents{
		Agent{
			AgentId: 1,
		},
		Agent{
			AgentId: 2,
		},
	}
	res, err := client.GetAgents()
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgentsError(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetAgents()
	teardown()
	assert.Error(t, err)
}

func TestClient_GetAgentJsonError(t *testing.T) {
	out := `{"agents":[{"agentId":4492,agentName":"Dallas, TX (Trial)","agentType":"Cloud","countryId":"US","ipAddresses":["104.130.154.136","104.130.156.108","104.130.141.203","104.130.155.161"],"location":"Dallas Area"}]}`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAgents()
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}
