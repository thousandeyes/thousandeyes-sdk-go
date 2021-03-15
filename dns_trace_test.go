package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetDNSTrace(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-trace","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","dnsTransportProtocol":  "UDP","apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-trace/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := DNSTrace{
		TestID:               122621,
		Enabled:              1,
		CreatedBy:            "William Fleming (wfleming@grumpysysadm.com)",
		CreatedDate:          "2020-02-06 15:28:07",
		SavedEvent:           0,
		AlertsEnabled:        1,
		TestName:             "test123",
		Type:                 "dns-trace",
		Interval:             300,
		Domain:               "webex.com",
		DNSTransportProtocol: "UDP",
		Agents: []Agent{
			{
				AgentID:     48620,
				AgentType:   "Cloud",
				AgentName:   "Seattle, WA (Trial) - IPv6",
				CountryID:   "US",
				IPAddresses: []string{"135.84.184.153"},
				Location:    "Seattle Area",
				Network:     "Astute Hosting Inc. (AS 54527)",
				Prefix:      "135.84.184.0/22",
			},
		},
		SharedWithAccounts: []SharedWithAccount{
			{
				AID:              176592,
				AccountGroupName: "Cloudreach",
			},
		},
		APILinks: APILinks{
			{
				Href: "https://api.thousandeyes.com/v6/tests/1226221",
				Rel:  "self",
			},
			{
				Href: "https://api.thousandeyes.com/v6/web/dns-trace/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/metrics/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/path-vis/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/bgp-metrics/1226221",
				Rel:  "data",
			},
		},
	}

	res, err := client.GetDNSTrace(122621)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetDNSTraceJsonError(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07",createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-trace","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","dnsTransportProtocol":  "UDP"}]"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-trace/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetDNSTrace(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'c' looking for beginning of object key string")
}

func TestClient_CreateDNSTrace(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-trace","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","dnsTransportProtocol":  "UDP","apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-trace/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-trace/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	// Define expected values from the API (based on the JSON we print out above)
	expected := DNSTrace{

		TestID:               122621,
		Enabled:              1,
		CreatedBy:            "William Fleming (wfleming@grumpysysadm.com)",
		CreatedDate:          "2020-02-06 15:28:07",
		SavedEvent:           0,
		TestName:             "test123",
		Type:                 "dns-trace",
		Interval:             300,
		AlertsEnabled:        1,
		Domain:               "webex.com",
		DNSTransportProtocol: "UDP",
		Agents: []Agent{
			{
				AgentID:     48620,
				AgentType:   "Cloud",
				AgentName:   "Seattle, WA (Trial) - IPv6",
				CountryID:   "US",
				IPAddresses: []string{"135.84.184.153"},
				Location:    "Seattle Area",
				Network:     "Astute Hosting Inc. (AS 54527)",
				Prefix:      "135.84.184.0/22",
			},
		},
		SharedWithAccounts: []SharedWithAccount{
			{
				AID:              176592,
				AccountGroupName: "Cloudreach",
			},
		},

		APILinks: APILinks{
			{
				Href: "https://api.thousandeyes.com/v6/tests/1226221",
				Rel:  "self",
			},
			{
				Href: "https://api.thousandeyes.com/v6/web/dns-trace/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/metrics/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/path-vis/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/bgp-metrics/1226221",
				Rel:  "data",
			},
		},
	}
	create := DNSTrace{
		TestName: "test1",
		Domain:   "1.1.1.1",
		Interval: 300,
	}
	res, err := client.CreateDNSTrace(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteDNSTrace(t *testing.T) {
	setup()

	mux.HandleFunc("/tests/dns-trace/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteDNSTrace(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_AddDnstraceAlertRule(t *testing.T) {
	test := DNSTrace{TestName: "test", AlertRules: []AlertRule{}}
	expected := DNSTrace{TestName: "test", AlertRules: []AlertRule{{RuleID: 1}}}
	test.AddAlertRule(1)
	assert.Equal(t, expected, test)
}

func TestClient_UpdateDNSTrace(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"dns-trace","domain":"webex.com" }]}`
	mux.HandleFunc("/tests/dns-trace/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	dnsS := DNSTrace{Domain: "webex.com"}
	res, err := client.UpdateDNSTrace(id, dnsS)
	if err != nil {
		t.Fatal(err)
	}
	expected := DNSTrace{TestID: 1, TestName: "test123", Type: "dns-trace", Domain: "webex.com"}
	assert.Equal(t, &expected, res)

}

func TestDNSTrace_AddAgent(t *testing.T) {
	test := DNSTrace{TestName: "test", Agents: Agents{}}
	expected := DNSTrace{TestName: "test", Agents: []Agent{{AgentID: 1}}}
	test.AddAgent(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetDNSTraceError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-trace/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetDNSTrace(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_GetDNSTraceStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"dns-trace"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetDNSTrace(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_CreateDNSTraceStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-trace/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreateDNSTrace(DNSTrace{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_UpdateDNSTraceStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-trace/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdateDNSTrace(1, DNSTrace{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_DeleteDNSTraceStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-trace/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeleteDNSTrace(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}
