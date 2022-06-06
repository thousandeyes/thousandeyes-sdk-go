package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetDNSSec(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-dnssec","interval":300,"alertsEnabled":1,"liveShare":0,"agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","dnsTransportProtocol":  "UDP","apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-dnssec/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	expected := DNSSec{
		TestID:        Int64(122621),
		Enabled:       Bool(true),
		CreatedBy:     String("ThousandEyes SRE (test.example@thousandeyes.com)"),
		CreatedDate:   String("2020-02-06 15:28:07"),
		SavedEvent:    Bool(false),
		AlertsEnabled: Bool(true),
		TestName:      String("test123"),
		Type:          String("dns-dnssec"),
		Interval:      Int(300),
		LiveShare:     Bool(false),
		Domain:        String("webex.com"),
		Agents: &[]Agent{
			{
				AgentID:     Int(48620),
				AgentType:   String("Cloud"),
				AgentName:   String("Seattle, WA (Trial) - IPv6"),
				CountryID:   String("US"),
				IPAddresses: &[]string{"135.84.184.153"},
				Location:    String("Seattle Area"),
				Network:     String("Astute Hosting Inc. (AS 54527)"),
				Prefix:      String("135.84.184.0/22"),
			},
		},
		SharedWithAccounts: &[]SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},
		APILinks: &[]APILink{
			{
				Href: String("https://api.thousandeyes.com/v6/tests/1226221"),
				Rel:  String("self"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/web/dns-dnssec/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/metrics/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/path-vis/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"),
				Rel:  String("data"),
			},
		},
	}

	res, err := client.GetDNSSec(122621)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetDNSSecJsonError(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07",createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-dnssec","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","dnsTransportProtocol":  "UDP"}]"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-dnssec/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetDNSSec(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'c' looking for beginning of object key string")
}

func TestClient_CreateDNSSec(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-dnssec","interval":300,"alertsEnabled":1,"liveShare":0,"agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-dnssec/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-dnssec/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	expected := DNSSec{
		TestID:        Int64(122621),
		Enabled:       Bool(true),
		CreatedBy:     String("ThousandEyes SRE (test.example@thousandeyes.com)"),
		CreatedDate:   String("2020-02-06 15:28:07"),
		SavedEvent:    Bool(false),
		TestName:      String("test123"),
		Type:          String("dns-dnssec"),
		Interval:      Int(300),
		LiveShare:     Bool(false),
		AlertsEnabled: Bool(true),
		Domain:        String("webex.com"),
		Agents: &[]Agent{
			{
				AgentID:     Int(48620),
				AgentType:   String("Cloud"),
				AgentName:   String("Seattle, WA (Trial) - IPv6"),
				CountryID:   String("US"),
				IPAddresses: &[]string{"135.84.184.153"},
				Location:    String("Seattle Area"),
				Network:     String("Astute Hosting Inc. (AS 54527)"),
				Prefix:      String("135.84.184.0/22"),
			},
		},
		SharedWithAccounts: &[]SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},

		APILinks: &[]APILink{
			{
				Href: String("https://api.thousandeyes.com/v6/tests/1226221"),
				Rel:  String("self"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/web/dns-dnssec/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/metrics/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/path-vis/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"),
				Rel:  String("data"),
			},
		},
	}
	create := DNSSec{
		TestName: String("test123"),
		Domain:   String("webex.com"),
		Interval: Int(300),
	}
	res, err := client.CreateDNSSec(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteDNSSec(t *testing.T) {
	setup()
	mux.HandleFunc("/tests/dns-dnssec/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteDNSSec(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_AddDNSSecAlertRule(t *testing.T) {
	test := DNSSec{TestName: String("test"), AlertRules: &[]AlertRule{}}
	expected := DNSSec{TestName: String("test"), AlertRules: &[]AlertRule{{RuleID: Int(1)}}}
	test.AddAlertRule(1)
	assert.Equal(t, expected, test)
}

func TestClient_UpdateDNSSec(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"dns-dnssec","domain":"webex.com" }]}`
	mux.HandleFunc("/tests/dns-dnssec/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	dnsp := DNSSec{Domain: String("webex.com")}
	res, err := client.UpdateDNSSec(id, dnsp)
	if err != nil {
		t.Fatal(err)
	}
	expected := DNSSec{TestID: Int64(1), TestName: String("test123"), Type: String("dns-dnssec"), Domain: String("webex.com")}
	assert.Equal(t, &expected, res)

}

func TestDNSSec_AddAgent(t *testing.T) {
	test := DNSSec{TestName: String("test"), Agents: &[]Agent{}}
	expected := DNSSec{TestName: String("test"), Agents: &[]Agent{{AgentID: Int(1)}}}
	test.AddAgent(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetDNSSecError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-dnssec/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetDNSSec(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_GetDNSSecStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"dnsp-trace"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetDNSSec(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_CreateDNSSecStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-dnssec/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.CreateDNSSec(DNSSec{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_UpdateDNSSecStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-dnssec/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.UpdateDNSSec(1, DNSSec{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_DeleteDNSSecStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-dnssec/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	err := client.DeleteDNSSec(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}
