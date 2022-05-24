package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAccountGroups(t *testing.T) {
	out := `{"accountGroups":[{"accountGroupName":"Test Account", "aid":1}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}

	mux.HandleFunc("/account-groups.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := []SharedWithAccount{
		{
			AccountGroupName: String("Test Account"),
			AID:              Int(1),
		},
	}

	res, err := client.GetAccountGroups()
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAccountGroupsAlertError(t *testing.T) {

	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}

	mux.HandleFunc("/account-groups.json", func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetAccountGroups()
	teardown()
	assert.Error(t, err)

}

func TestClient_GetAccountGroupsJsonError(t *testing.T) {
	out := `{"accountGroups":[{accountGroupName":"Test Account", "aid":1}]}`

	setup()

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/account-groups.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAccountGroups()
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}
