package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// DNSTrace - DNS trace test
type DNSTrace struct {
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
	Agents               *[]Agent `json:"agents,omitempty"`
	DNSTransportProtocol *string  `json:"dnsTransportProtocol,omitempty"`
	Domain               *string  `json:"domain,omitempty"`
	Interval             *int     `json:"interval,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t DNSTrace) MarshalJSON() ([]byte, error) {
	type aliasTest DNSTrace

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *DNSTrace) UnmarshalJSON(data []byte) error {
	type aliasTest DNSTrace
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAgent - Add agent to DNS Trace test
func (t *DNSTrace) AddAgent(id int64) {
	agent := Agent{AgentID: Int64(id)}
	*t.Agents = append(*t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *DNSTrace) AddAlertRule(id int64) {
	alertRule := AlertRule{RuleID: Int64(id)}
	*t.AlertRules = append(*t.AlertRules, alertRule)
}

// GetDNSTrace - get dns trace test
func (c *Client) GetDNSTrace(id int) (*DNSTrace, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &DNSTrace{}, err
	}
	var target map[string][]DNSTrace
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateDNSTrace - Create dns trace test
func (c Client) CreateDNSTrace(t DNSTrace) (*DNSTrace, error) {
	resp, err := c.post("/tests/dns-trace/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSTrace
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteDNSTrace - delete dns trace test
func (c *Client) DeleteDNSTrace(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/dns-trace/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete dns trace test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateDNSTrace - update dns trace test
func (c *Client) UpdateDNSTrace(id int, t DNSTrace) (*DNSTrace, error) {
	resp, err := c.post(fmt.Sprintf("/tests/dns-trace/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSTrace
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
