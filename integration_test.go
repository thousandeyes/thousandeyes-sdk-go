package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetIntegrations(t *testing.T) {
	out := `{"integrations":{"thirdParty":[{"authMethod":"Auth Token","integrationId":"pgd-9999","integrationName":"Test PD Integration","integrationType":"PAGER_DUTY"}],"webhook":[{"authMethod":"Basic","integrationId":"wb-999","integrationName":"Test Webhook Integration","integrationType":"WEBHOOK","target":"https://thousandeyes.com/"}]}}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}

	mux.HandleFunc("/integrations.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := []Integration{
		{
			AuthMethod:      String("Auth Token"),
			IntegrationID:   String("pgd-9999"),
			IntegrationName: String("Test PD Integration"),
			IntegrationType: String("PAGER_DUTY"),
		},
		{
			AuthMethod:      String("Basic"),
			IntegrationID:   String("wb-999"),
			IntegrationName: String("Test Webhook Integration"),
			IntegrationType: String("WEBHOOK"),
			Target:          String("https://thousandeyes.com/"),
		},
	}

	res, err := client.GetIntegrations()
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetIntegrationsAlertError(t *testing.T) {

	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}

	mux.HandleFunc("/integrations.json", func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetIntegrations()
	teardown()
	assert.Error(t, err)

}

func TestClient_GetIntegrationsJsonError(t *testing.T) {
	out := ` { "bgpMonitors": [ {aonitorId":1, "monitorType": "bgp","monitorName": "test", "ipAddress": "1.2.3.4"} ] }`

	setup()

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/integrations.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetIntegrations()
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}
