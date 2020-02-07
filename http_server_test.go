package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_GetHttpServer(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"http-server","interval":300,"url":"https://test.com","protocol":"TCP","ipv6Policy":"USE_AGENT_POLICY","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBgp":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"followRedirects":1,"sslVersionId":0,"verifyCertificate":1,"useNtlm":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":64,"ipAddress":"2001:240:100:ff::2497:2","countryId":"JP","monitorName":"Tokyo-3","network":"IIJ Internet Initiative Japan Inc. (AS 2497)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}],"sslVersion":"Auto"}]}`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := HttpServer{
		TestId:                122621,
		Enabled:               1,
		CreatedBy:             "William Fleming (wfleming@grumpysysadm.com)",
		CreatedDate:           "2020-02-06 15:28:07",
		SavedEvent:            0,
		TestName:              "test123",
		Type:                  "http-server",
		Interval:              300,
		Url:                   "https://test.com",
		Protocol:              "TCP",
		NetworkMeasurements:   1,
		MtuMeasurements:       1,
		BandwidthMeasurements: 0,
		BgpMeasurements:       1,
		AlertsEnabled:         1,
		LiveShare:             0,
		HttpTimeLimit:         5,
		HttpTargetTime:        1000,
		HttpVersion:           2,
		FollowRedirects:       1,
		NumPathTraces:         3,
		SslVersionId:          0,
		VerifyCertificate:     1,
		UseNtlm:               0,
		AuthType:              "NONE",
		ContentRegex:          "",
		ProbeMode:             "AUTO",
		Agents: []Agent{
			{
				AgentId:     48620,
				AgentName:   "Seattle, WA (Trial) - IPv6",
				CountryId:   "US",
				IpAddresses: []string{"135.84.184.153"},
				Location:    "Seattle Area",
				Network:     "Astute Hosting Inc. (AS 54527)",
				Prefix:      "135.84.184.0/22",
			},
		},
		SharedWithAccounts: []AccountGroup{
			{
				Aid:  176592,
				Name: "Cloudreach",
			},
		},
		BgpMonitors: []Monitor{
			{
				MonitorId:   64,
				IpAddress:   "2001:240:100:ff::2497:2",
				CountryId:   "JP",
				MonitorName: "Tokyo-3",
				Network:     "IIJ Internet Initiative Japan Inc. (AS 2497)",
				MonitorType: "Public",
			},
		},
		SslVersion: "Auto",
		ApiLinks: []ApiLink{
			{
				Href: "https://api.thousandeyes.com/v6/tests/1226221",
				Rel:  "self",
			},
			{
				Href: "https://api.thousandeyes.com/v6/web/http-server/1226221",
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

	res, err := client.GetHttpServer(122621)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteHttpServer(t *testing.T) {
	setup()

	mux.HandleFunc("/tests/http-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteHttpServer(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateHttpServer(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"http-server","url":"https://test.com"}]}`
	mux.HandleFunc("/tests/http-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	httpS := HttpServer{Url: "https://test.com"}
	res, err := client.UpdateHttpServer(id, httpS)
	if err != nil {
		t.Fatal(err)
	}
	expected := HttpServer{TestId: 1, TestName: "test123", Type: "http-server", Url: "https://test.com"}
	assert.Equal(t, &expected, res)

}
