package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_GetUsers(t *testing.T) {
	setup()
	out := `{"users": [{"name": "ThousandEyes SRE", "email": "test.example@thousandeyes.com", "uid": 1}, {"name": "Test User 2", "email": "test.example@thousandeyes.com", "uid": 2}]}`
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
			Name:  String("ThousandEyes SRE"),
			Email: String("test.example@thousandeyes.com"),
			UID:   Int(1),
		},
		{
			Name:  String("Test User 2"),
			Email: String("test.example@thousandeyes.com"),
			UID:   Int(2),
		},
	}
	assert.Equal(t, &expected, res)
}

func TestClient_GetUser(t *testing.T) {
	setup()
	out := `{"users": [{"name": "ThousandEyes SRE", "email": "test.example@thousandeyes.com", "uid": 1}]}`
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
		Name:  String("ThousandEyes SRE"),
		Email: String("test.example@thousandeyes.com"),
		UID:   Int(1),
	}
	assert.Equal(t, &expected, res)
}

func TestClient_CreateUser(t *testing.T) {
	setup()
	out := `{"name": "ThousandEyes SRE", "email": "test.example@thousandeyes.com", "uid": 1}`
	mux.HandleFunc("/users/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	create := User{
		Name:  String("ThousandEyes SRE"),
		Email: String("test.example@thousandeyes.com"),
	}
	res, err := client.CreateUser(create)
	if err != nil {
		t.Fatal(err)
	}

	expected := User{
		Name:  String("ThousandEyes SRE"),
		Email: String("test.example@thousandeyes.com"),
		UID:   Int(1),
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
	out := `{"name": "ThousandEyes SRE", "email": "text.example@thousandeyes.com", "uid": 1}`
	mux.HandleFunc("/users/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	update := User{
		Email: String("text.example@thousandeyes.com"),
	}
	res, err := client.UpdateUser(1, update)
	if err != nil {
		t.Fatal(err)
	}

	expected := User{
		Name:  String("ThousandEyes SRE"),
		Email: String("text.example@thousandeyes.com"),
		UID:   Int(1),
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
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
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
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
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
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
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
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_GetUsersJsonError(t *testing.T) {
	out := `{"users": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/users.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetUsers()
	assert.Error(t, err)
	assert.EqualError(t, err, "could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GetUserJsonError(t *testing.T) {
	out := `{"users": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/users/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetUser(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_UpdateUsersJsonError(t *testing.T) {
	out := `{"users": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/users/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.UpdateUser(1, User{})
	assert.Error(t, err)
	assert.EqualError(t, err, "could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_CreateUsersJsonError(t *testing.T) {
	out := `{"users": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/users/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.CreateUser(User{})
	assert.Error(t, err)
	assert.EqualError(t, err, "could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}
