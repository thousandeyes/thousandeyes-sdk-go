package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// DNSSec - DNSSec test
type DNSSec struct {
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
	Agents   *[]Agent `json:"agents,omitempty"`
	Domain   *string  `json:"domain,omitempty"`
	Interval *int     `json:"interval,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t DNSSec) MarshalJSON() ([]byte, error) {
	type aliasTest DNSSec

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *DNSSec) UnmarshalJSON(data []byte) error {
	type aliasTest DNSSec
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAgent - Add agent to DNSSec test
func (t *DNSSec) AddAgent(id int64) {
	agent := Agent{AgentID: Int64(id)}
	*t.Agents = append(*t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *DNSSec) AddAlertRule(id int64) {
	alertRule := AlertRule{RuleID: Int64(id)}
	*t.AlertRules = append(*t.AlertRules, alertRule)
}

// GetDNSSec - get DNSSec test
func (c *Client) GetDNSSec(id int64) (*DNSSec, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &DNSSec{}, err
	}
	var target map[string][]DNSSec
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateDNSSec - Create DNSSec test
func (c Client) CreateDNSSec(t DNSSec) (*DNSSec, error) {
	resp, err := c.post("/tests/dns-dnssec/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create dns dnssec test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSSec
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// DeleteDNSSec - delete DNSSec test
func (c *Client) DeleteDNSSec(id int64) error {
	resp, err := c.post(fmt.Sprintf("/tests/dns-dnssec/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete dnsp domain test, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateDNSSec - update DNSSec test
func (c *Client) UpdateDNSSec(id int64, t DNSSec) (*DNSSec, error) {
	resp, err := c.post(fmt.Sprintf("/tests/dns-dnssec/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSSec
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
