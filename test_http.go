package go_thousandeyes

import (
	"errors"
	"fmt"
)

type HttpTest struct {
	Agents                Agents         `json:"agents,omitempty"`
	AlertsEnabled         string         `json:"alertsEnabled,omitempty"`
	AlertRules            []AlertRule    `json:"alertRules,omitempty"`
	ApiLinks              []interface{}  `json:"apiLinks,omitempty"`
	CreatedBy             string         `json:"createdBy,omitempty"`
	CreatedDate           string         `json:"createdDate,omitempty"`
	Description           string         `json:"description,omitempty"`
	Enabled               int            `json:"enabled,omitempty"`
	Groups                []GroupLabels  `json:"groups,omitempty"`
	LiveShare             int            `json:"liveShare,omitempty"`
	ModifiedBy            string         `json:"modifiedBy,omitempty"`
	ModifiedDate          string         `json:"modifiedDate,omitempty"`
	SavedEvent            int            `json:"savedEvent,omitempty"`
	SharedWithAccounts    []AccountGroup `json:"sharedWithAccounts,omitempty"`
	TestId                int            `json:"testId,omitempty"`
	TestName              string         `json:"testName,omitempty"`
	Type                  string         `json:"type,omitempty"`
	AuthType              string         `json:"authType,omitempty"`
	BandwidthMeasurements string         `json:"bandwidthMeasurements,omitempty"`
	BgpMeasurements       int            `json:"bgpMeasurements,omitempty"`
	BgpMonitors           Monitors       `json:"bgpMonitors,omitempty"`
	ClientCertificate     string         `json:"clientCertificate,omitempty"`
	ContentRegex          string         `json:"contentRegex,omitempty"`
	DesiredStatusCode     string         `json:"desiredStatusCode,omitempty"`
	DownloadLimit         string         `json:"downloadLimit,omitempty"`
	DnsOverride           string         `json:"dnsOverride,omitempty"`
	FollowRedirects       int            `json:"followRedirects,omitempty"`
	Headers               []string       `json:"headers,omitempty"`
	HttpVersion           int            `json:"httpVersion,omitempty"`
	HttpTargetTime        int            `json:"httpTargetTime,omitempty"`
	HttpTimeLimit         int            `json:"httpTimeLimit,omitempty"`
	Interval              int            `json:"interval,omitempty"`
	MtuMeasurements       int            `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int            `json:"networkMeasurements,omitempty"`
	NumPathTraces         int            `json:"numPathTraces,omitempty"`
	Password              string         `json:"password,omitempty"`
	PostBody              string         `json:"postBody,omitempty"`
	ProbeMode             string         `json:"probeMode,omitempty"`
	Protocol              string         `json:"protocol,omitempty"`
	SslVersion            string         `json:"sslVersion,omitempty"`
	SslVersionId          int            `json:"sslVersionId,omitempty"`
	Url                   string         `json:"url,omitempty"`
	UseNtlm               int            `json:"useNtlm,omitempty"`
	UserAgent             string         `json:"userAgent,omitempty"`
	Username              string         `json:"username,omitempty"`
	VerifyCertificate     int            `json:"verifyCertificate,omitempty"`
}

func (t *HttpTest) AddAgent(id int) {
	agent := Agent{AgentId: id}
	t.Agents = append(t.Agents, agent)
}

func (c *Client) GetHttpTest(id string) (*HttpTest, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%s", id))
	if err != nil {
		return &HttpTest{}, err
	}
	var target map[string]HttpTest
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	rootNode := "test"
	t, nodeOK := target[rootNode]
	if !nodeOK {
		return nil, fmt.Errorf("JSON response does not have %s field", rootNode)
	}
	return &t, nil
}

func (c Client) CreateHttpTest(t HttpTest) (*HttpTest, error) {
	res, err := c.post("/tests/http-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if res.StatusCode != 201 {
		return &t, errors.New(fmt.Sprintf("failed to create test, response code %d", res.StatusCode))
	}
	return &t, nil
}

func (c *Client) DeleteHttpTest(id string) error {
	_, err := c.delete(fmt.Sprintf("/tests/http-server/%s/delete", id))
	return err
}

func (c *Client) UpdateHttpTest(t HttpTest) error {
	return nil
}
