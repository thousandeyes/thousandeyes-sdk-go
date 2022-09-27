package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// AgentAgent - test
type AgentAgent struct {
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
	Agents                 *[]Agent      `json:"agents,omitempty"`
	BGPMeasurements        *bool         `json:"bgpMeasurements,omitempty" te:"int-bool"`
	BGPMonitors            *[]BGPMonitor `json:"bgpMonitors,omitempty"`
	Direction              *string       `json:"direction,omitempty"`
	DSCP                   *string       `json:"dscp,omitempty"`
	DSCPID                 *int64        `json:"dscpId"`
	Interval               *int          `json:"interval,omitempty"`
	MSS                    *int          `json:"mss,omitempty"`
	NetworkMeasurements    *bool         `json:"networkMeasurements,omitempty" te:"int-bool"`
	MTUMeasurements        *bool         `json:"mtuMeasurements,omitempty" te:"int-bool"`
	NumPathTraces          *int          `json:"numPathTraces,omitempty"`
	PathTraceMode          *string       `json:"pathTraceMode,omitempty"`
	Port                   *int          `json:"port,omitempty"`
	Protocol               *string       `json:"protocol,omitempty"`
	TargetAgentID          *int64        `json:"targetAgentId,omitempty"`
	ThroughputDuration     *int          `json:"throughputDuration,omitempty"`
	ThroughputMeasurements *bool         `json:"throughputMeasurements,omitempty" te:"int-bool"`
	ThroughputRate         *int          `json:"throughputRate,omitempty"`
	UsePublicBGP           *bool         `json:"usePublicBgp,omitempty" te:"int-bool"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t AgentAgent) MarshalJSON() ([]byte, error) {
	type aliasTest AgentAgent

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *AgentAgent) UnmarshalJSON(data []byte) error {
	type aliasTest AgentAgent
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAgent - Adds an agent to agent test
func (t *AgentAgent) AddAgent(id int64) {
	agent := Agent{AgentID: Int64(id)}
	*t.Agents = append(*t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *AgentAgent) AddAlertRule(id int64) {
	alertRule := AlertRule{RuleID: Int64(id)}
	*t.AlertRules = append(*t.AlertRules, alertRule)
}

// GetAgentAgent - Get an agent to agent test
func (c *Client) GetAgentAgent(id int64) (*AgentAgent, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &AgentAgent{}, err
	}
	var target map[string][]AgentAgent
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateAgentAgent - Create an agent to agent test
func (c Client) CreateAgentAgent(t AgentAgent) (*AgentAgent, error) {
	resp, err := c.post("/tests/agent-to-agent/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create agent test, response code %d", resp.StatusCode)
	}
	var target map[string][]AgentAgent
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// DeleteAgentAgent - delete agent to agent test
func (c *Client) DeleteAgentAgent(id int64) error {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-agent/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete agent test, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateAgentAgent - update agent to agent test
func (c *Client) UpdateAgentAgent(id int64, t AgentAgent) (*AgentAgent, error) {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-agent/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update agent test, response code %d", resp.StatusCode)
	}
	var target map[string][]AgentAgent
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
