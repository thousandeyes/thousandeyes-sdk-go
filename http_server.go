package thousandeyes

import (
	"fmt"
)

type HttpServerResponse struct {
	tests []HttpServer
}

type HttpServer struct {
	Agents                Agents         `json:"agents,omitempty"`
	AlertsEnabled         int            `json:"alertsEnabled,omitempty"`
	AlertRules            []AlertRule    `json:"alertRules,omitempty"`
	ApiLinks              []ApiLink      `json:"apiLinks,omitempty"`
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
	BandwidthMeasurements int            `json:"bandwidthMeasurements,omitempty"`
	BgpMeasurements       int            `json:"bgpMeasurements,omitempty"`
	BgpMonitors           []Monitor      `json:"bgpMonitors,omitempty"`
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

func (t *HttpServer) AddAgent(id int) {
	agent := Agent{AgentId: id}
	t.Agents = append(t.Agents, agent)
}

func (c *Client) GetHttpServer(id int) (*HttpServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &HttpServer{}, err
	}
	var target map[string][]HttpServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

func (c Client) CreateHttpServer(t HttpServer) (*HttpServer, error) {
	resp, err := c.post("/tests/http-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create http server, response code %d", resp.StatusCode)
	}
	var target map[string][]HttpServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

func (c *Client) DeleteHttpServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/http-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete http server, response code %d", resp.StatusCode)
	}
	return nil
}

func (c *Client) UpdateHttpServer(id int, t HttpServer) (*HttpServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/http-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update http server, response code %d", resp.StatusCode)
	}
	var target map[string][]HttpServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
