package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_CreateAgentServer(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","Port": 8090}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentServer{
		TestId:      1,
		TestName:    "test",
		CreatedDate: "2020-02-06 15:28:07",
		CreatedBy:   "William Fleming (wfleming@grumpysysadm.com)",
		Port:        8090,
	}
	create := AgentServer{
		TestName: "test",
		Port:     8090,
	}
	res, err := client.CreateAgentServer(create)
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgentServer(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","Port": 8090}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentServer{
		TestId:      1,
		TestName:    "test",
		CreatedDate: "2020-02-06 15:28:07",
		CreatedBy:   "William Fleming (wfleming@grumpysysadm.com)",
		Port:        8090,
	}
	res, err := client.GetAgentServer(1)
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteAgentServer(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/tests/agent-to-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteAgentServer(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateAgentServer(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)", "Port": 8090}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentServer{
		TestId:      1,
		TestName:    "test",
		CreatedDate: "2020-02-06 15:28:07",
		CreatedBy:   "William Fleming (wfleming@grumpysysadm.com)",
		Port:        8090,
	}
	update := AgentServer{
		Port: 8090,
	}
	res, err := client.UpdateAgentServer(1, update)
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgentServerError(t *testing.T) {
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetAgentServer(1)
	assert.Error(t, err)
}
