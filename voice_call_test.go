package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetVoiceCall(t *testing.T) {
	out := `{"test":[{"createdDate":"2018-11-03 19:09:42","modifiedDate":"2019-02-06 01:09:56","createdBy":"ThousandEyes (support@thousandeyes.com)","modifiedBy":"ThousandEyes (support@thousandeyes.com)","enabled":1,"savedEvent":0,"testId":814641,"testName":"Voice Call - AWS SIP server","interval":120,"server":"18.234.180.66:5060","bgpMeasurements":1,"usePublicBGP":1,"duration":5,"codec":"G.711 @ 64 Kbps","codecId":0,"dscpId":46,"jitterBuffer":40,"sipTimeLimit":5,"alertsEnabled":0,"liveShare":0,"targetAgentId":69,"numPathTraces":3,"sourceSIPCredentials":{"credentialsId":48162,"user":"1006","sipRegistrar":"18.234.180.66","sipProxy":"","authUser":"1006","port":5060,"protocol":"UDP"},"targetSIPCredentials":{"credentialsId":48165,"user":"1005","sipRegistrar":"18.234.180.66","sipProxy":"","authUser":"1005","port":5060,"protocol":"UDP"},"sipTargetTime":1000,"dscp":"EF (DSCP 46)","apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/814641"},{"rel":"data","href":"https://api.thousandeyes.com/v6/voice/sip-server/814641"},{"rel":"data","href":"https://api.thousandeyes.com/v6/voice/rtp-stream/814641"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/814641"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := VoiceCall{
		TestID:          814641,
		Enabled:         1,
		CreatedBy:       "ThousandEyes (support@thousandeyes.com)",
		CreatedDate:     "2018-11-03 19:09:42",
		SavedEvent:      0,
		AlertsEnabled:   0,
		TestName:        "Voice Call - AWS SIP server",
		Interval:        120,
		Duration:        5,
		JitterBuffer:    40,
		SIPTargetTime:   1000,
		SIPTimeLimit:    5,
		ModifiedDate:    "2019-02-06 01:09:56",
		ModifiedBy:      "ThousandEyes (support@thousandeyes.com)",
		TargetAgentID:   69,
		Codec:           "G.711 @ 64 Kbps",
		BGPMeasurements: 1,
		UsePublicBGP:    1,
		NumPathTraces:   3,
		DSCP:            "EF (DSCP 46)",
		DSCPID:          46,
		TargetSIPCredentials: SIPAuthData{
			Protocol:     "UDP",
			AuthUser:     "1005",
			Password:     "",
			Port:         5060,
			SIPRegistrar: "18.234.180.66",
			User:         "1005",
		},
		SourceSIPCredentials: SIPAuthData{
			Protocol:     "UDP",
			AuthUser:     "1006",
			Password:     "",
			Port:         5060,
			SIPRegistrar: "18.234.180.66",
			User:         "1006",
		},
		APILinks: APILinks{

			{
				Rel:  "self",
				Href: "https://api.thousandeyes.com/v6/tests/814641",
			},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/voice/sip-server/814641",
			},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/voice/rtp-stream/814641",
			},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/net/bgp-metrics/814641",
			},
		},
	}

	res, err := client.GetVoiceCall(122621)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetVoiceCallJsonError(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07",createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"sip-server","interval":300,"alertsEnabled":1,"liveShare":0,"agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","dnsTransportProtocol":  "UDP"}]"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-trace/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetVoiceCall(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'c' looking for beginning of object key string")
}

func TestClient_CreateVoiceCall(t *testing.T) {
	out := `{"test":[{"createdDate":"2018-11-03 19:09:42","modifiedDate":"2019-02-06 01:09:56","createdBy":"ThousandEyes (support@thousandeyes.com)","duration" : 5,"modifiedBy":"ThousandEyes (support@thousandeyes.com)","enabled":1,"savedEvent":0,"testId":814641,"testName":"Voice Call - AWS SIP server","interval":120,"server":"18.234.180.66:5060","bgpMeasurements":1,"usePublicBGP":1,"codec":"G.711 @ 64 Kbps","codecId":0,"dscpId":46,"alertsEnabled":0,"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/814641"},{"rel":"data","href":"https://api.thousandeyes.com/v6/voice/sip-server/814641"},{"rel":"data","href":"https://api.thousandeyes.com/v6/voice/rtp-stream/814641"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/814641"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/voice-call/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := VoiceCall{
		TestID:          814641,
		Enabled:         1,
		ModifiedBy:      "ThousandEyes (support@thousandeyes.com)",
		ModifiedDate:    "2019-02-06 01:09:56",
		CreatedBy:       "ThousandEyes (support@thousandeyes.com)",
		CreatedDate:     "2018-11-03 19:09:42",
		SavedEvent:      0,
		TestName:        "Voice Call - AWS SIP server",
		Interval:        120,
		AlertsEnabled:   0,
		DSCPID:          46,
		Duration:        5,
		BGPMeasurements: 1,
		UsePublicBGP:    1,
		NumPathTraces:   3,
		Codec:           "G.711 @ 64 Kbps",
		APILinks: APILinks{
			{
				Href: "https://api.thousandeyes.com/v6/tests/814641",
				Rel:  "self",
			},
			{
				Href: "https://api.thousandeyes.com/v6/voice/sip-server/814641",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/voice/rtp-stream/814641",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/bgp-metrics/814641",
				Rel:  "data",
			},
		},
	}
	create := VoiceCall{
		TestName: "test1",
		DSCPID:   46,
		Duration: 5,
		Interval: 120,
		Codec:    "G.711 @ 64 Kbps",
	}
	res, err := client.CreateVoiceCall(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteVoiceCall(t *testing.T) {
	setup()

	mux.HandleFunc("/tests/voice-call/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteVoiceCall(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateVoiceCall(t *testing.T) {
	setup()
	out := `{"test":[{"testId": 1234,"testName":"Voice Call - AWS SIP server","interval":120,"server":"18.234.180.66:5060","codec":"G.711 @ 64 Kbps","codecId":0,"jitterBuffer":40,"alertsEnabled":0}]}`
	mux.HandleFunc("/tests/voice-call/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	sipS := VoiceCall{
		TestName:     "Voice Call - AWS SIP server",
		CodecID:      0,
		Codec:        "G.711 @ 64 Kbps",
		JitterBuffer: 40,
	}
	res, err := client.UpdateVoiceCall(id, sipS)
	if err != nil {
		t.Fatal(err)
	}
	expected := VoiceCall{Interval: 120, TestID: 1234, Codec: "G.711 @ 64 Kbps", TestName: "Voice Call - AWS SIP server", CodecID: 0, JitterBuffer: 40}
	assert.Equal(t, &expected, res)

}

func TestVoiceCall_AddAgent(t *testing.T) {
	test := VoiceCall{TestName: "test", Agents: Agents{}}
	expected := VoiceCall{TestName: "test", Agents: []Agent{{AgentID: 1}}}
	test.AddAgent(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetVoiceCallError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/voice-call/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetVoiceCall(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_GetVoiceCallStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"sip-server"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetVoiceCall(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_CreateVoiceCallStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/voice-call/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreateVoiceCall(VoiceCall{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_UpdateVoiceCallStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/voice-call/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdateVoiceCall(1, VoiceCall{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_DeleteVoiceCallStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/voice-call/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeleteVoiceCall(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}
