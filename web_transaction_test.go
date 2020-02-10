package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_CreateWebTransaction(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","transactionScript":"script here"}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/web-transaction/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := WebTransaction{
		TestId:            1,
		TestName:          "test",
		CreatedDate:       "2020-02-06 15:28:07",
		CreatedBy:         "William Fleming (wfleming@grumpysysadm.com)",
		TransactionScript: "script here",
	}
	create := WebTransaction{
		TestName:          "test",
		TransactionScript: "script here",
	}
	res, err := client.CreateWebTransaction(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetWebTransaction(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","transactionScript":"script here"}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := WebTransaction{
		TestId:            1,
		TestName:          "test",
		CreatedDate:       "2020-02-06 15:28:07",
		CreatedBy:         "William Fleming (wfleming@grumpysysadm.com)",
		TransactionScript: "script here",
	}
	res, err := client.GetWebTransaction(1)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}
