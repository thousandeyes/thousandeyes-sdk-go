package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_GetUsers(t *testing.T) {
	setup()
	out := `{"users": [{"name": "William Fleming", "email": "wfleming@grumpysysadm.com", "uid": 1}, {"name": "Test User 2", "email": "wfleming@grumpysysadm.com", "uid": 2}]}`
	mux.HandleFunc("/users.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}

	res, err := client.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
	expected := []User{
		{
			Name:  "William Fleming",
			Email: "wfleming@grumpysysadm.com",
			UID:   1,
		},
		{
			Name:  "Test User 2",
			Email: "wfleming@grumpysysadm.com",
			UID:   2,
		},
	}
	assert.Equal(t, &expected, res)
}

func TestClient_GetUser(t *testing.T) {
	setup()
	out := `{"users": [{"name": "William Fleming", "email": "wfleming@grumpysysadm.com", "uid": 1}]}`
	mux.HandleFunc("/users/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}

	res, err := client.GetUser(1)
	if err != nil {
		t.Fatal(err)
	}
	expected := User{
		Name:  "William Fleming",
		Email: "wfleming@grumpysysadm.com",
		UID:   1,
	}
	assert.Equal(t, &expected, res)
}

func TestClient_CreateUser(t *testing.T) {
	setup()
	out := `{"name": "William Fleming", "email": "wfleming@grumpysysadm.com", "uid": 1}`
	mux.HandleFunc("/users/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	create := User{
		Name:  "William Fleming",
		Email: "wfleming@grumpysysadm.com",
	}
	res, err := client.CreateUser(create)
	if err != nil {
		t.Fatal(err)
	}

	expected := User{
		Name:  "William Fleming",
		Email: "wfleming@grumpysysadm.com",
		UID:   1,
	}
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteUser(t *testing.T) {
	setup()
	mux.HandleFunc("/users/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}

	_ = client.DeleteUser(1)
}

func TestClient_UpdateUser(t *testing.T) {
	setup()
	out := `{"name": "William Fleming", "email": "william@grumpysysadm.com", "uid": 1}`
	mux.HandleFunc("/users/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	update := User{
		Email: "william@grumpysysadm.com",
	}
	res, err := client.UpdateUser(1, update)
	if err != nil {
		t.Fatal(err)
	}

	expected := User{
		Name:  "William Fleming",
		Email: "william@grumpysysadm.com",
		UID:   1,
	}
	assert.Equal(t, &expected, res)
}

func TestClient_GetUserStatusCode(t *testing.T) {
	setup()
	out := `{"alertRules":[{"ruleId":1,"ruleName":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/users.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetUsers()
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_CreateUserStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/users/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.CreateUser(User{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_UpdateUserStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/users/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.UpdateUser(1, User{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_DeleteUserStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/users/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	err := client.DeleteUser(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}
