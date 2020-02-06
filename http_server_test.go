package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHttpServer_Get(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (william.fleming@cloudreach.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"http-server","interval":300,"url":"https://dashboards.coqa.cloudreach.com","protocol":"TCP","ipv6Policy":"USE_AGENT_POLICY","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBgp":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"followRedirects":1,"sslVersionId":0,"verifyCertificate":1,"useNtlm":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":64,"ipAddress":"2001:240:100:ff::2497:2","countryId":"JP","monitorName":"Tokyo-3","network":"IIJ Internet Initiative Japan Inc. (AS 2497)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}],"sslVersion":"Auto"}]}`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := HttpServer{
		TestId:                122621,
		Enabled:               1,
		CreatedBy:             "William Fleming (william.fleming@cloudreach.com)",
		CreatedDate:           "2020-02-06 15:28:07",
		SavedEvent:            0,
		TestName:              "test123",
		Type:                  "http-server",
		Interval:              300,
		Url:                   "https://dashboards.coqa.cloudreach.com",
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
				AgentId:   48620,
				AgentName: "Seattle, WA (Trial) - IPv6",
				//AgentType:   "Cloud",
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
