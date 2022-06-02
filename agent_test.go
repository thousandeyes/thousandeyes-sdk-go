package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAgents(t *testing.T) {
	out := `{"agents":[{"agentId": 1, "enabled": 1}, {"agentId": 2, "enabled": 0}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := Agents{
		Agent{
			AgentID: Int(1),
			Enabled: Bool(true),
		},
		Agent{
			AgentID: Int(2),
			Enabled: Bool(false),
		},
	}
	res, err := client.GetAgents()
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgent(t *testing.T) {
	out := `{"agents":[{"agentId": 1, "enabled": 1}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	expected := Agent{AgentID: Int(1), Enabled: Bool(true)}
	res, err := client.GetAgent(1)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgentsError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetAgents()
	teardown()
	assert.Error(t, err)
}

func TestClient_GetAgentsJsonError(t *testing.T) {
	out := `{"agents":[{"agentId":4492,agentName":"Dallas, TX (Trial)","agentType":"Cloud","countryId":"US","ipAddresses":["104.130.154.136","104.130.156.108","104.130.141.203","104.130.155.161"],"location":"Dallas Area"}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAgents()
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}

func TestClient_GetAgentJsonError(t *testing.T) {
	out := `{"agents":[{"agentId":4492,agentName":"Dallas, TX (Trial)","agentType":"Cloud","countryId":"US","ipAddresses":["104.130.154.136","104.130.156.108","104.130.141.203","104.130.155.161"],"location":"Dallas Area"}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAgent(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}

func TestClient_GetAgentsStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetAgents()
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_GetAgentStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetAgent(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_AddAgentToClusterStatusCode(t *testing.T) {
	setup()
	out := `{"agents": [{"agentId": 1, "agentName": "test", "clusterMembers": [{"memberId": 80001, "name": "test"}]}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1/add-to-cluster.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.AddAgentsToCluster(1, []int{8001})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_RemoveAgentToClusterStatusCode(t *testing.T) {
	setup()
	out := `{"agents": [{"agentId": 1, "agentName": "test", "clusterMembers": [{"memberId": 80001, "name": "test"}]}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1/remove-from-cluster.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.RemoveAgentsFromCluster(1, []int{8001})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_AddAgentToClusterJsonError(t *testing.T) {
	out := `{"agents": ["agentId": 1, "agentName": "test", "clusterMembers": [{"memberId": 80001, "name": "test"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1/add-to-cluster.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.AddAgentsToCluster(1, []int{8001})
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character ':' after array element")
}

func TestClient_RemoveAgentFromClusterJsonError(t *testing.T) {
	out := `{"agents": ["agentId": 1, "agentName": "test", "clusterMembers": [{"memberId": 80001, "name": "test"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1/remove-from-cluster.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.RemoveAgentsFromCluster(1, []int{8001})
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character ':' after array element")
}

func TestClient_RemoveAgentFromCluster(t *testing.T) {
	out := `{"agents": [{"agentId": 1, "agentName": "test", "clusterMembers": [{"memberId": 80002, "name": "test"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1/remove-from-cluster.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})
	res, _ := client.RemoveAgentsFromCluster(1, []int{8001})
	exp := []Agent{
		{
			AgentID:   Int(1),
			AgentName: String("test"),
			ClusterMembers: &[]ClusterMember{
				{
					MemberID: Int(80002),
					Name:     String("test"),
				},
			},
		},
	}
	assert.Equal(t, res, &exp)
}

func TestClient_AddAgentToCluster(t *testing.T) {
	out := `{"agents": [{"agentId": 1, "agentName": "test", "clusterMembers": [{"memberId": 80002, "name": "test"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1/add-to-cluster.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})
	res, _ := client.AddAgentsToCluster(1, []int{8002})
	exp := []Agent{
		{
			AgentID:   Int(1),
			AgentName: String("test"),
			ClusterMembers: &[]ClusterMember{
				{
					MemberID: Int(80002),
					Name:     String("test"),
				},
			},
		},
	}
	assert.Equal(t, res, &exp)
}
