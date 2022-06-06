package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_AddAgentSeverAlertRule(t *testing.T) {
	test := AgentServer{TestName: String("test"), AlertRules: &[]AlertRule{}}
	expected := AgentServer{TestName: String("test"), AlertRules: &[]AlertRule{{RuleID: Int(1)}}}
	test.AddAlertRule(1)
	assert.Equal(t, expected, test)
}

func TestClient_AgentServerAddAgent(t *testing.T) {
	test := AgentServer{TestName: String("test"), Agents: &[]Agent{}}
	expected := AgentServer{TestName: String("test"), Agents: &[]Agent{{AgentID: Int(1)}}}
	test.AddAgent(1)
	assert.Equal(t, expected, test)
}

func TestClient_CreateAgentServer(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)","server":"thousandeyes.com:8090"}]}`
	setup()
	defer teardown()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentServer{
		TestID:      Int64(1),
		TestName:    String("test"),
		CreatedDate: String("2020-02-06 15:28:07"),
		CreatedBy:   String("ThousandEyes SRE (test.example@thousandeyes.com)"),
		Port:        Int(8090),
		Server:      String("thousandeyes.com"),
	}
	create := AgentServer{
		TestName: String("test"),
		Port:     Int(8090),
		Server:   String("thousandeyes.com"),
	}
	res, err := client.CreateAgentServer(create)
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_AddAgentServerAlertRule(t *testing.T) {
	test := AgentServer{TestName: String("test"), AlertRules: &[]AlertRule{}}
	expected := AgentServer{TestName: String("test"), AlertRules: &[]AlertRule{{RuleID: Int(1)}}}
	test.AddAlertRule(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetAgentServerJsonError(t *testing.T) {
	out := `{"test":[test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAgentServer(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GetAgentServer(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)","Port": 8090}]}`
	setup()
	defer teardown()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentServer{
		TestID:      Int64(1),
		TestName:    String("test"),
		CreatedDate: String("2020-02-06 15:28:07"),
		CreatedBy:   String("ThousandEyes SRE (test.example@thousandeyes.com)"),
		Port:        Int(8090),
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

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteAgentServer(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateAgentServer(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)", "Port": 8090}]}`
	setup()
	defer teardown()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	expected := AgentServer{
		TestID:      Int64(1),
		TestName:    String("test"),
		CreatedDate: String("2020-02-06 15:28:07"),
		CreatedBy:   String("ThousandEyes SRE (test.example@thousandeyes.com)"),
		Port:        Int(8090),
	}
	update := AgentServer{
		Port: Int(8090),
	}
	res, err := client.UpdateAgentServer(1, update)
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgentServerError(t *testing.T) {
	setup()
	defer teardown()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetAgentServer(1)
	assert.Error(t, err)
}

func TestClient_GetAgentServerStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetPageLoad(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_CreateAgentServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.CreateAgentServer(AgentServer{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_UpdateAgentServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.UpdateAgentServer(1, AgentServer{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_DeleteAgentServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/agent-to-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	err := client.DeleteAgentServer(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestExtractPort(t *testing.T) {
	test := AgentServer{
		Agents: &[]Agent{
			{
				AgentID: Int(75),
			},
		},
		Interval: Int(3600),
		Server:   String("foo.com:8888"),
	}
	result, err := extractPort(test)
	if err != nil {
		assert.Error(t, err)
	}
	test.Server = String("foo.com")
	test.Port = Int(8888)
	assert.Equal(t, test, result)
}
