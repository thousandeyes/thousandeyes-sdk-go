package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetGroupLabels(t *testing.T) {
	out := `{"groups" : [ {"groupId":1, "type" : "tests" , "name": "exampleName" }]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := GroupLabels{
		GroupLabel{GroupID: Int64(1), Type: String("tests"), Name: String("exampleName")},
	}

	res, err := client.GetGroupLabels()
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)

}

func TestClient_GetGroupLabelsByType(t *testing.T) {
	out := `{"groups" : [ {"groupId":1, "type" : "tests", "name": "test-agent", "builtin": 0 }]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/tests.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := GroupLabels{
		GroupLabel{GroupID: Int64(1), Builtin: Bool(false), Type: String("tests"), Name: String("test-agent")},
	}

	res, err := client.GetGroupLabelsByType("tests")
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetGroupLabel(t *testing.T) {
	out := `{
		"groups" : [
			{
				"groupId" : 222, "type" : "tests", "name" : "test-agent", "builtin": 0
			}
		]
	}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/222.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := GroupLabel{
		GroupID: Int64(222), Builtin: Bool(false), Type: String("tests"), Name: String("test-agent"),
	}

	res, err := client.GetGroupLabel(222)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetGroupLabelsError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetGroupLabels()
	teardown()
	assert.Error(t, err)
}
func TestClient_CreateGroupLabelError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/tests/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.CreateGroupLabel(GroupLabel{})
	teardown()
	assert.Error(t, err)
}

func TestClient_GetGroupLabelError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetGroupLabel(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_GetGroupLabelByTypeError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/tests.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetGroupLabelsByType("tests")
	teardown()
	assert.Error(t, err)
}

func TestClient_DeleteGroupLabel(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/groups/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := int64(1)
	err := client.DeleteGroupLabel(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateGroupLabel(t *testing.T) {
	setup()
	out := `{ "groups" : [ { "groupId" : 222, "type" : "tests", "name" : "test-agent" } ] }`

	mux.HandleFunc("/groups/222/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := int64(222)
	u := GroupLabel{Type: String("tests")}
	res, err := client.UpdateGroupLabel(id, u)
	if err != nil {
		t.Fatal(err)
	}
	expected := GroupLabels{GroupLabel{GroupID: Int64(222), Type: String("tests"), Name: String("test-agent")}}
	assert.Equal(t, &expected, res)
}

func TestClient_CreateGroupLabel(t *testing.T) {
	setup()
	out := `{"groups" : [ {"groupId":1, "name": "test"}]}`
	mux.HandleFunc("/groups/tests/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	u := GroupLabel{Name: String("test"), Type: String("tests"), GroupID: Int64(1)}
	res, err := client.CreateGroupLabel(u)
	if err != nil {
		t.Fatal(err)
	}
	expected := GroupLabel{GroupID: Int64(1), Name: String("test")}
	assert.Equal(t, &expected, res)
}

func TestClient_GroupLabelsJsonError(t *testing.T) {
	out := `{"groups": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetGroupLabels()
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GroupLabelsByTypeJsonError(t *testing.T) {
	out := `{"groups": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetGroupLabelsByType("agents")
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GroupLabelError(t *testing.T) {
	out := `{"groups": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetGroupLabel(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GetGroupLabelsStatusCode(t *testing.T) {
	setup()
	out := `{"groups":[{"groupId":1,"name":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetGroupLabels()
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_CreateGroupLabelJsonError(t *testing.T) {
	out := `{"groups": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/tests/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))

	})
	_, err := client.CreateGroupLabel(GroupLabel{Type: String("tests")})
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_DeleteGroupLabelStatusCodeGood(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteGroupLabel(1)
	teardown()
	assert.Nil(t, err)
}

func TestClient_DeleteGroupLabelStatusCodeBad(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	err := client.DeleteGroupLabel(1)
	teardown()
	assert.NotNil(t, err)

}
func TestClient_CreateGroupLabelStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/tests/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreateGroupLabel(GroupLabel{Type: String("tests")})
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_UpdateGroupLabelStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdateGroupLabel(1, GroupLabel{})
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_DeleteGroupLabelStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/groups/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeleteGroupLabel(1)
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}
