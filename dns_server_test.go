package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetDNSServer(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","numPathTraces":3,"pathTraceMode": "classic", "enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-server","interval":300,"protocol":"UDP","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBGP":1,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","recursiveQueries":0,"agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":64,"ipAddress":"2001:240:100:ff::2497:2","countryId":"JP","monitorName":"Tokyo-3","network":"IIJ Internet Initiative Japan Inc. (AS 2497)","monitorType":"Public"}],"numPathTraces":3,"domain": "webex.com","dnsTransportProtocol":  "UDP", "dnsServers" : [{"serverId": 123,"serverName":"1.1.1.1"}],"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := DNSServer{

		TestID:                Int64(122621),
		Enabled:               Int(1),
		CreatedBy:             String("William Fleming (wfleming@grumpysysadm.com)"),
		CreatedDate:           String("2020-02-06 15:28:07"),
		SavedEvent:            Int(0),
		TestName:              String("test123"),
		Type:                  String("dns-server"),
		Interval:              Int(300),
		LiveShare:             Int(0),
		Protocol:              String("UDP"),
		NetworkMeasurements:   Int(1),
		MTUMeasurements:       Int(1),
		BandwidthMeasurements: Int(0),
		NumPathTraces:         Int(3),
		PathTraceMode:         String("classic"),
		AlertsEnabled:         Int(1),
		RecursiveQueries:      Int(0),
		BGPMeasurements:       Int(1),
		UsePublicBGP:          Int(1),
		Domain:                String("webex.com"),
		ProbeMode:             String("AUTO"),
		DNSTransportProtocol:  String("UDP"),
		Agents: []Agent{
			{
				AgentID:     Int(48620),
				AgentType:   String("Cloud"),
				AgentName:   String("Seattle, WA (Trial) - IPv6"),
				CountryID:   String("US"),
				IPAddresses: []string{"135.84.184.153"},
				Location:    String("Seattle Area"),
				Network:     String("Astute Hosting Inc. (AS 54527)"),
				Prefix:      String("135.84.184.0/22"),
			},
		},
		SharedWithAccounts: []SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},
		DNSServers: []Server{
			{
				ServerID:   Int(123),
				ServerName: String("1.1.1.1"),
			},
		},
		BGPMonitors: []BGPMonitor{
			{
				MonitorID:   Int(64),
				IPAddress:   String("2001:240:100:ff::2497:2"),
				MonitorName: String("Tokyo-3"),
				Network:     String("IIJ Internet Initiative Japan Inc. (AS 2497)"),
				MonitorType: String("Public"),
			},
		},
		APILinks: APILinks{
			{
				Href: String("https://api.thousandeyes.com/v6/tests/1226221"),
				Rel:  String("self"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/web/dns-server/1226221"),
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

	res, err := client.GetDNSServer(122621)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_AddDnsserverAlertRule(t *testing.T) {
	test := DNSServer{TestName: String("test"), AlertRules: []AlertRule{}}
	expected := DNSServer{TestName: String("test"), AlertRules: []AlertRule{{RuleID: Int(1)}}}
	test.AddAlertRule(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetDNSServerJsonError(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07",createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-server","interval":300,"protocol":"UDP","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBGP":1,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":64,"ipAddress":"2001:240:100:ff::2497:2","countryId":"JP","monitorName":"Tokyo-3","network":"IIJ Internet Initiative Japan Inc. (AS 2497)","monitorType":"Public"}],"numPathTraces":3,"domain": "webex.com","dnsTransportProtocol":  "UDP", "dnsServers" : [{"serverId": 123,"serverName":"1.1.1.1"}]}]"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetDNSServer(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'c' looking for beginning of object key string")
}

func TestClient_CreateDNSServer(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","numPathTraces": 3,"enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"dns-server","interval":300,"protocol":"UDP","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBGP":1,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","recursiveQueries":0,"agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":64,"ipAddress":"2001:240:100:ff::2497:2","countryId":"JP","monitorName":"Tokyo-3","network":"IIJ Internet Initiative Japan Inc. (AS 2497)","monitorType":"Public"}],"numPathTraces":3,"domain": "webex.com","dnsTransportProtocol":  "UDP", "dnsServers" : [{"serverId": 123,"serverName":"1.1.1.1"}],"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	// Define expected values from the API (based on the JSON we print out above)
	expected := DNSServer{

		TestID:                Int64(122621),
		Enabled:               Int(1),
		CreatedBy:             String("William Fleming (wfleming@grumpysysadm.com)"),
		CreatedDate:           String("2020-02-06 15:28:07"),
		SavedEvent:            Int(0),
		TestName:              String("test123"),
		Type:                  String("dns-server"),
		Interval:              Int(300),
		LiveShare:             Int(0),
		Protocol:              String("UDP"),
		NetworkMeasurements:   Int(1),
		MTUMeasurements:       Int(1),
		BandwidthMeasurements: Int(0),
		NumPathTraces:         Int(3),
		RecursiveQueries:      Int(0),
		AlertsEnabled:         Int(1),
		BGPMeasurements:       Int(1),
		UsePublicBGP:          Int(1),
		Domain:                String("webex.com"),
		ProbeMode:             String("AUTO"),
		DNSTransportProtocol:  String("UDP"),
		Agents: []Agent{
			{
				AgentID:     Int(48620),
				AgentType:   String("Cloud"),
				AgentName:   String("Seattle, WA (Trial) - IPv6"),
				CountryID:   String("US"),
				IPAddresses: []string{"135.84.184.153"},
				Location:    String("Seattle Area"),
				Network:     String("Astute Hosting Inc. (AS 54527)"),
				Prefix:      String("135.84.184.0/22"),
			},
		},
		SharedWithAccounts: []SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},
		DNSServers: []Server{
			{
				ServerID:   Int(123),
				ServerName: String("1.1.1.1"),
			},
		},
		BGPMonitors: []BGPMonitor{
			{
				MonitorID:   Int(64),
				IPAddress:   String("2001:240:100:ff::2497:2"),
				MonitorName: String("Tokyo-3"),
				Network:     String("IIJ Internet Initiative Japan Inc. (AS 2497)"),
				MonitorType: String("Public"),
			},
		},
		APILinks: APILinks{
			{
				Href: String("https://api.thousandeyes.com/v6/tests/1226221"),
				Rel:  String("self"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/web/dns-server/1226221"),
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
	create := DNSServer{
		TestName:      String("test1"),
		Domain:        String("1.1.1.1"),
		Interval:      Int(300),
		NumPathTraces: Int(3),
	}
	res, err := client.CreateDNSServer(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteDNSServer(t *testing.T) {
	setup()

	mux.HandleFunc("/tests/dns-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteDNSServer(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateDNSServer(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"dns-server","domain":"webex.com" }]}`
	mux.HandleFunc("/tests/dns-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	dnsS := DNSServer{Domain: String("webex.com")}
	res, err := client.UpdateDNSServer(id, dnsS)
	if err != nil {
		t.Fatal(err)
	}
	expected := DNSServer{TestID: Int64(1), TestName: String("test123"), Type: String("dns-server"), Domain: String("webex.com")}
	assert.Equal(t, &expected, res)

}

func TestDNSServer_AddAgent(t *testing.T) {
	test := DNSServer{TestName: String("test"), Agents: Agents{}}
	expected := DNSServer{TestName: String("test"), Agents: []Agent{{AgentID: Int(1)}}}
	test.AddAgent(Int(1))
	assert.Equal(t, expected, test)
}

func TestClient_GetDNSServerError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-server/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetDNSServer(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_GetDNSServerStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"dns-server"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetDNSServer(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_CreateDNSServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreateDNSServer(DNSServer{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_UpdateDNSServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdateDNSServer(1, DNSServer{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_DeleteDNSServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/dns-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeleteDNSServer(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}
