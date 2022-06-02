package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetFTPServer(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"ftp-server","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"url": "webex.com","apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/ftp-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := FTPServer{
		TestID:        Int64(122621),
		Enabled:       Bool(true),
		CreatedBy:     String("William Fleming (wfleming@grumpysysadm.com)"),
		CreatedDate:   String("2020-02-06 15:28:07"),
		SavedEvent:    Bool(false),
		AlertsEnabled: Bool(true),
		TestName:      String("test123"),
		Type:          String("ftp-server"),
		LiveShare:     Bool(false),
		Interval:      Int(300),
		URL:           String("webex.com"),
		ProbeMode:     String("AUTO"),
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
		APILinks: &[]APILink{
			{
				Href: String("https://api.thousandeyes.com/v6/tests/1226221"),
				Rel:  String("self"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/web/ftp-server/1226221"),
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
		SharedWithAccounts: &[]SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},
	}

	res, err := client.GetFTPServer(122621)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetFTPServerJsonError(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07",createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"ftp-server","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","dnsTransportProtocol":  "UDP"}]"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-trace/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetFTPServer(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'c' looking for beginning of object key string")
}

func TestClient_CreateFTPServer(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"ftp-server","interval":300,"alertsEnabled":1,"liveShare":0,"protocol": "TCP","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"url": "webex.com","apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/ftp-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/ftp-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := FTPServer{
		TestID:        Int64(122621),
		Enabled:       Bool(true),
		CreatedBy:     String("William Fleming (wfleming@grumpysysadm.com)"),
		CreatedDate:   String("2020-02-06 15:28:07"),
		SavedEvent:    Bool(false),
		TestName:      String("test123"),
		Type:          String("ftp-server"),
		AlertsEnabled: Bool(true),
		Interval:      Int(300),
		LiveShare:     Bool(false),
		URL:           String("webex.com"),
		Protocol:      String("TCP"),
		ProbeMode:     String("AUTO"),
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
				Href: String("https://api.thousandeyes.com/v6/web/ftp-server/1226221"),
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
	create := FTPServer{
		TestName:  String("test123"),
		URL:       String("webex.com"),
		Protocol:  String("TCP"),
		ProbeMode: String("AUTO"),
	}
	res, err := client.CreateFTPServer(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteFTPServer(t *testing.T) {
	setup()

	mux.HandleFunc("/tests/ftp-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteFTPServer(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_AddFTPServerAlertRule(t *testing.T) {
	test := FTPServer{TestName: String("test"), AlertRules: &[]AlertRule{}}
	expected := FTPServer{TestName: String("test"), AlertRules: &[]AlertRule{{RuleID: Int(1)}}}
	test.AddAlertRule(1)
	assert.Equal(t, expected, test)
}

func TestClient_UpdateFTPServer(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"ftp-server","url":"webex.com" }]}`
	mux.HandleFunc("/tests/ftp-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	ftpServer := FTPServer{URL: String("webex.com")}
	res, err := client.UpdateFTPServer(id, ftpServer)
	if err != nil {
		t.Fatal(err)
	}
	expected := FTPServer{TestID: Int64(1), TestName: String("test123"), Type: String("ftp-server"), URL: String("webex.com")}
	assert.Equal(t, &expected, res)

}

func TestFTPServer_AddAgent(t *testing.T) {
	test := FTPServer{TestName: String("test"), Agents: &[]Agent{}}
	expected := FTPServer{TestName: String("test"), Agents: &[]Agent{{AgentID: Int(1)}}}
	test.AddAgent(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetFTPServerError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/ftp-server/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetFTPServer(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_GetFTPServerStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"ftp-server"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetFTPServer(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_CreateFTPServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/ftp-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.CreateFTPServer(FTPServer{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_UpdateFTPServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/ftp-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.UpdateFTPServer(1, FTPServer{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_DeleteFTPServerCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/ftp-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	err := client.DeleteFTPServer(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}
