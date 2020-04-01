package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_CreateAgentAgent(t *testing.T) {
	out := `{"test": [{"testID":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","throughputMeasurements": 1, "Port": 8090}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-agent/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentAgent{
		TestID:                 1,
		TestName:               "test",
		CreatedDate:            "2020-02-06 15:28:07",
		CreatedBy:              "William Fleming (wfleming@grumpysysadm.com)",
		Port:                   8090,
		ThroughputMeasurements: 1,
	}
	create := AgentAgent{
		TestName: "test",
		Port:     8090,
	}
	res, err := client.CreateAgentAgent(create)
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgentAgentJsonError(t *testing.T) {
	out := `{"test":[test]}`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAgentAgent(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GetAgentAgent(t *testing.T) {
	out := `{"test": [{"testID":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","Port": 8090, "throughputMeasurements" : 1, "throughputDuration":10000}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentAgent{
		TestID:                 1,
		TestName:               "test",
		CreatedDate:            "2020-02-06 15:28:07",
		CreatedBy:              "William Fleming (wfleming@grumpysysadm.com)",
		Port:                   8090,
		ThroughputDuration:     10000,
		ThroughputMeasurements: 1,
	}
	res, err := client.GetAgentAgent(1)
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteAgentAgent(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/tests/agent-to-agent/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteAgentAgent(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateAgentAgent(t *testing.T) {
	out := `{"test": [{"testID":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)", "Port": 8090}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-agent/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentAgent{
		TestID:      1,
		TestName:    "test",
		CreatedDate: "2020-02-06 15:28:07",
		CreatedBy:   "William Fleming (wfleming@grumpysysadm.com)",
		Port:        8090,
	}
	update := AgentAgent{
		Port: 8090,
	}
	res, err := client.UpdateAgentAgent(1, update)
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgentAgentError(t *testing.T) {
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-agent/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetAgentAgent(1)
	assert.Error(t, err)
}

func TestClient_GetAgentAgentStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testID":1,"testName":"test123"}]}`
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetPageLoad(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_CreateAgentAgentStatusCode(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-agent/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreateAgentAgent(AgentAgent{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_UpdateAgentAgentStatusCode(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-agent/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdateAgentAgent(1, AgentAgent{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_DeleteAgentAgentStatusCode(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-agent/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeleteAgentAgent(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}
