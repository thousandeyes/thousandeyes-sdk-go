package thousandeyes

import (
	"fmt"
)

// PageLoad - a page log struct
type PageLoad struct {
	Agents                Agents        `json:"agents,omitempty"`
	AlertsEnabled         int           `json:"alertsEnabled,omitempty"`
	AlertRules            AlertRules    `json:"alertRules,omitempty"`
	APILinks              APILinks      `json:"apiLinks,omitempty"`
	CreatedBy             string        `json:"createdBy,omitempty"`
	CreatedDate           string        `json:"createdDate,omitempty"`
	Description           string        `json:"description,omitempty"`
	Enabled               int           `json:"enabled,omitempty"`
	Groups                []GroupLabel  `json:"groups,omitempty"`
	LiveShare             int           `json:"liveShare,omitempty"`
	ModifiedBy            string        `json:"modifiedBy,omitempty"`
	ModifiedDate          string        `json:"modifiedDate,omitempty"`
	SavedEvent            int           `json:"savedEvent,omitempty"`
	SharedWithAccounts    AccountGroups `json:"sharedWithAccounts,omitempty"`
	TestID                int           `json:"testId,omitempty"`
	TestName              string        `json:"testName,omitempty"`
	Type                  string        `json:"type,omitempty"`
	AuthType              string        `json:"authType,omitempty"`
	BandwidthMeasurements int           `json:"bandwidthMeasurements,omitempty"`
	BgpMeasurements       int           `json:"bgpMeasurements,omitempty"`
	BgpMonitors           BGPMonitors   `json:"bgpMonitors,omitempty"`
	HTTPInterval          int           `json:"httpInterval,omitempty"`
	HTTPVersion           int           `json:"httpVersion,omitempty"`
	HTTPTargetTime        int           `json:"httpTargetTime,omitempty"`
	HTTPTimeLimit         int           `json:"httpTimeLimit,omitempty"`
	IncludeHeaders        int           `json:"includeHeaders,omitempty"`
	Interval              int           `json:"interval,omitempty"`
	MtuMeasurements       int           `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int           `json:"networkMeasurements,omitempty"`
	NumPathTraces         int           `json:"numPathTraces,omitempty"`
	PageLoadTargetTime    int           `json:"pageLoadTargetTime,omitempty"`
	PageLoadTimeLimit     int           `json:"pageLoadTimeLimit,omitempty"`
	Password              string        `json:"password,omitempty"`
	ProbeMode             string        `json:"probeMode,omitempty"`
	Protocol              string        `json:"protocol,omitempty"`
	SslVersion            string        `json:"sslVersion,omitempty"`
	SslVersionID          int           `json:"sslVersionId,omitempty"`
	URL                   string        `json:"url,omitempty"`
	UseNtlm               int           `json:"useNtlm,omitempty"`
	UserAgent             string        `json:"userAgent,omitempty"`
	Username              string        `json:"username,omitempty"`
	VerifyCertificate     int           `json:"verifyCertificate,omitempty"`
}

// AddAgent  - add an aget
func (t *PageLoad) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

//GetPageLoad - get page load test
func (c *Client) GetPageLoad(id int) (*PageLoad, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &PageLoad{}, err
	}
	var target map[string][]PageLoad
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreatePageLoad - create pager load test
func (c Client) CreatePageLoad(t PageLoad) (*PageLoad, error) {
	resp, err := c.post("/tests/page-load/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create test, response code %d", resp.StatusCode)
	}
	var target map[string][]PageLoad
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// DeletePageLoad - Delete page load tes
func (c *Client) DeletePageLoad(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/page-load/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete page load, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdatePageLoad - Upload page load
func (c *Client) UpdatePageLoad(id int, t PageLoad) (*PageLoad, error) {
	resp, err := c.post(fmt.Sprintf("/tests/page-load/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]PageLoad
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
