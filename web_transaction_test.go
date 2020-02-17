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
	mux.HandleFunc("/tests/web-transactions/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
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
		_, _ = w.Write([]byte(out))
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

func TestClient_DeleteWebTransaction(t *testing.T) {
	setup()

	mux.HandleFunc("/tests/web-transactions/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteWebTransaction(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateWebTransaction(t *testing.T) {
	out := `{"test": [{"testId":1,"testName":"test","createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","transactionScript":"new script here"}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/web-transactions/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := WebTransaction{
		TestId:            1,
		TestName:          "test",
		CreatedDate:       "2020-02-06 15:28:07",
		CreatedBy:         "William Fleming (wfleming@grumpysysadm.com)",
		TransactionScript: "new script here",
	}
	update := WebTransaction{
		TransactionScript: "new script here",
	}
	res, err := client.UpdateWebTransaction(1, update)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetWebTransactionError(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/web-transactions/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetWebTransaction(1)
	teardown()
	assert.Error(t, err)
}
