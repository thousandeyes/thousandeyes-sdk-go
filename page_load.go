package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// PageLoad - a page log struct
type PageLoad struct {
	// Common test fields
	AlertsEnabled      *bool                `json:"alertsEnabled,omitempty" te:"int-bool"`
	AlertRules         *[]AlertRule         `json:"alertRules,omitempty"`
	APILinks           *[]APILink           `json:"apiLinks,omitempty"`
	CreatedBy          *string              `json:"createdBy,omitempty"`
	CreatedDate        *string              `json:"createdDate,omitempty"`
	Description        *string              `json:"description,omitempty"`
	Enabled            *bool                `json:"enabled,omitempty" te:"int-bool"`
	Groups             *[]GroupLabel        `json:"groups,omitempty"`
	ModifiedBy         *string              `json:"modifiedBy,omitempty"`
	ModifiedDate       *string              `json:"modifiedDate,omitempty"`
	SavedEvent         *bool                `json:"savedEvent,omitempty" te:"int-bool"`
	SharedWithAccounts *[]SharedWithAccount `json:"sharedWithAccounts,omitempty"`
	TestID             *int64               `json:"testId,omitempty"`
	TestName           *string              `json:"testName,omitempty"`
	Type               *string              `json:"type,omitempty"`
	LiveShare          *bool                `json:"liveShare,omitempty" te:"int-bool"`

	// Fields unique to this test
	Agents                *[]Agent       `json:"agents,omitempty"`
	AuthType              *string        `json:"authType,omitempty"`
	BandwidthMeasurements *bool          `json:"bandwidthMeasurements,omitempty" te:"int-bool"`
	BGPMeasurements       *bool          `json:"bgpMeasurements,omitempty" te:"int-bool"`
	BGPMonitors           *[]BGPMonitor  `json:"bgpMonitors,omitempty"`
	ContentRegex          *string        `json:"contentRegex,omitempty"`
	CustomHeaders         *CustomHeaders `json:"customHeaders,omitempty"`
	FollowRedirects       *bool          `json:"followRedirects,omitempty" te:"int-bool"`
	HTTPInterval          *int           `json:"httpInterval,omitempty"`
	HTTPTargetTime        *int           `json:"httpTargetTime,omitempty"`
	HTTPTimeLimit         *int           `json:"httpTimeLimit,omitempty"`
	HTTPVersion           *int           `json:"httpVersion,omitempty"`
	IncludeHeaders        *bool          `json:"includeHeaders,omitempty" te:"int-bool"`
	Interval              *int           `json:"interval,omitempty"`
	MTUMeasurements       *bool          `json:"mtuMeasurements,omitempty" te:"int-bool"`
	NetworkMeasurements   *bool          `json:"networkMeasurements,omitempty" te:"int-bool"`
	NumPathTraces         *int           `json:"numPathTraces,omitempty"`
	PageLoadTargetTime    *int           `json:"pageLoadTargetTime,omitempty"`
	PageLoadTimeLimit     *int           `json:"pageLoadTimeLimit,omitempty"`
	Password              *string        `json:"password,omitempty"`
	PathTraceMode         *string        `json:"pathTraceMode,omitempty"`
	ProbeMode             *string        `json:"probeMode,omitempty"`
	Protocol              *string        `json:"protocol,omitempty"`
	SSLVersion            *string        `json:"sslVersion,omitempty"`
	SSLVersionID          *int64         `json:"sslVersionId,omitempty"`
	Subinterval           *int           `json:"subinterval,omitempty"`
	URL                   *string        `json:"url,omitempty"`
	UseNTLM               *bool          `json:"useNtlm,omitempty" te:"int-bool"`
	UsePublicBGP          *bool          `json:"usePublicBgp,omitempty" te:"int-bool"`
	UserAgent             *string        `json:"userAgent,omitempty"`
	Username              *string        `json:"username,omitempty"`
	VerifyCertificate     *bool          `json:"verifyCertificate,omitempty" te:"int-bool"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t PageLoad) MarshalJSON() ([]byte, error) {
	type aliasTest PageLoad

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *PageLoad) UnmarshalJSON(data []byte) error {
	type aliasTest PageLoad
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAgent  - add an aget
func (t *PageLoad) AddAgent(id int64) {
	agent := Agent{AgentID: Int64(id)}
	*t.Agents = append(*t.Agents, agent)
}

//GetPageLoad - get page load test
func (c *Client) GetPageLoad(id int64) (*PageLoad, error) {
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
func (c *Client) DeletePageLoad(id int64) error {
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
func (c *Client) UpdatePageLoad(id int64, t PageLoad) (*PageLoad, error) {
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
