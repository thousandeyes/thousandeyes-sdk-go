package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_CreatePageLoad(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 19:15:36","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":1226422,"testName":"test1","type":"page-load","interval":300,"httpInterval":300,"url":"https://test.com","protocol":"TCP","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBGP":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"pageLoadTimeLimit":10,"pageLoadTargetTime":6,"followRedirects":1,"includeHeaders":1,"sslVersionId":0,"verifyCertificate":1,"useNtlm":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":62,"ipAddress":"2001:1890:111d:1::63","countryId":"US","monitorName":"New York, NY-6","network":"AT&T Services, Inc. (AS 7018)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/page-load/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226422"}],"sslVersion":"Auto"}]}`
	setup()
	defer teardown()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := PageLoad{
		CreatedDate:           String("2020-02-06 19:15:36"),
		CreatedBy:             String("William Fleming (wfleming@grumpysysadm.com)"),
		Enabled:               Bool(true),
		SavedEvent:            Bool(false),
		TestID:                Int64(1226422),
		TestName:              String("test1"),
		Type:                  String("page-load"),
		Interval:              Int(300),
		HTTPInterval:          Int(300),
		URL:                   String("https://test.com"),
		Protocol:              String("TCP"),
		FollowRedirects:       Bool(true),
		NetworkMeasurements:   Bool(true),
		MTUMeasurements:       Bool(true),
		BandwidthMeasurements: Bool(false),
		BGPMeasurements:       Bool(true),
		UsePublicBGP:          Bool(true),
		AlertsEnabled:         Bool(true),
		LiveShare:             Bool(false),
		HTTPTimeLimit:         Int(5),
		HTTPTargetTime:        Int(1000),
		HTTPVersion:           Int(2),
		PageLoadTimeLimit:     Int(10),
		PageLoadTargetTime:    Int(6),
		IncludeHeaders:        Bool(true),
		SSLVersionID:          Int(0),
		VerifyCertificate:     Bool(true),
		UseNTLM:               Bool(false),
		AuthType:              String("NONE"),
		ProbeMode:             String("AUTO"),
		ContentRegex:          String(""),
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
		BGPMonitors: []BGPMonitor{
			{
				MonitorID:   Int(62),
				IPAddress:   String("2001:1890:111d:1::63"),
				MonitorName: String("New York, NY-6"),
				Network:     String("AT&T Services, Inc. (AS 7018)"),
				MonitorType: String("Public"),
			},
		},
		NumPathTraces: Int(3),
		APILinks: APILinks{
			{
				Rel:  String("self"),
				Href: String("https://api.thousandeyes.com/v6/tests/1226422"),
			},
			{
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/web/http-server/1226422"),
			}, {
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/web/page-load/1226422")},
			{
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/net/metrics/1226422"),
			},
			{
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/net/path-vis/1226422"),
			}, {
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/net/bgp-metrics/1226422"),
			},
		},
		SSLVersion: String("Auto"),
	}
	create := PageLoad{
		TestName:     String("test1"),
		URL:          String("https://test.com"),
		Interval:     Int(300),
		HTTPInterval: Int(300),
	}
	res, err := client.CreatePageLoad(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetPageLoad(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 19:15:36","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":1226422,"testName":"test1","type":"page-load","interval":300,"httpInterval":300,"url":"https://test.com","protocol":"TCP","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBGP":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"pageLoadTimeLimit":10,"pageLoadTargetTime":6,"followRedirects":1,"includeHeaders":1,"sslVersionId":0,"verifyCertificate":1,"useNtlm":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":62,"ipAddress":"2001:1890:111d:1::63","countryId":"US","monitorName":"New York, NY-6","network":"AT&T Services, Inc. (AS 7018)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/page-load/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226422"}],"sslVersion":"Auto"}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1226422.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := PageLoad{
		CreatedDate:           String("2020-02-06 19:15:36"),
		CreatedBy:             String("William Fleming (wfleming@grumpysysadm.com)"),
		Enabled:               Bool(true),
		SavedEvent:            Bool(false),
		TestID:                Int64(1226422),
		TestName:              String("test1"),
		Type:                  String("page-load"),
		Interval:              Int(300),
		HTTPInterval:          Int(300),
		URL:                   String("https://test.com"),
		Protocol:              String("TCP"),
		FollowRedirects:       Bool(true),
		NetworkMeasurements:   Bool(true),
		MTUMeasurements:       Bool(true),
		BandwidthMeasurements: Bool(false),
		BGPMeasurements:       Bool(true),
		UsePublicBGP:          Bool(true),
		AlertsEnabled:         Bool(true),
		LiveShare:             Bool(false),
		HTTPTimeLimit:         Int(5),
		HTTPTargetTime:        Int(1000),
		HTTPVersion:           Int(2),
		PageLoadTimeLimit:     Int(10),
		PageLoadTargetTime:    Int(6),
		IncludeHeaders:        Bool(true),
		SSLVersionID:          Int(0),
		VerifyCertificate:     Bool(true),
		UseNTLM:               Bool(false),
		AuthType:              String("NONE"),
		ProbeMode:             String("AUTO"),
		ContentRegex:          String(""),
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
		BGPMonitors: []BGPMonitor{
			{
				MonitorID:   Int(62),
				IPAddress:   String("2001:1890:111d:1::63"),
				MonitorName: String("New York, NY-6"),
				Network:     String("AT&T Services, Inc. (AS 7018)"),
				MonitorType: String("Public"),
			},
		},
		NumPathTraces: Int(3),
		APILinks: APILinks{
			{
				Rel:  String("self"),
				Href: String("https://api.thousandeyes.com/v6/tests/1226422"),
			},
			{
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/web/http-server/1226422"),
			}, {
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/web/page-load/1226422")},
			{
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/net/metrics/1226422"),
			},
			{
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/net/path-vis/1226422"),
			}, {
				Rel:  String("data"),
				Href: String("https://api.thousandeyes.com/v6/net/bgp-metrics/1226422"),
			},
		},
		SSLVersion: String("Auto"),
	}

	res, err := client.GetPageLoad(1226422)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeletePageLoad(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/tests/page-load/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeletePageLoad(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdatePageLoad(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"page-load","url":"https://test.com"}]}`
	mux.HandleFunc("/tests/page-load/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	httpS := PageLoad{URL: String("https://test.com")}
	res, err := client.UpdatePageLoad(id, httpS)
	if err != nil {
		t.Fatal(err)
	}
	expected := PageLoad{TestID: Int64(1), TestName: String("test123"), Type: String("page-load"), URL: String("https://test.com")}
	assert.Equal(t, &expected, res)

}

func TestPageLoad_AddAgent(t *testing.T) {
	test := PageLoad{TestName: String("test"), Agents: Agents{}}
	expected := PageLoad{TestName: String("test"), Agents: []Agent{{AgentID: Int(1)}}}
	test.AddAgent(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetPageLoadError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetPageLoad(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_PageLoadJsonError(t *testing.T) {
	out := `{"test": [test]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetPageLoad(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GetPageLoadStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetPageLoad(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_CreatePageLoadStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreatePageLoad(PageLoad{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_UpdatePageLoadStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdatePageLoad(1, PageLoad{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}

func TestClient_DeletePageLoadStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeletePageLoad(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{<nil>}")
}
