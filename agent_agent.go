package thousandeyes

import (
	"fmt"
)

// AgentAgent - test
type AgentAgent struct {
	Agents                 []Agent        `json:"agents,omitempty"`
	AlertsEnabled          int            `json:"alertsEnabled,omitempty"`
	AlertRules             []AlertRule    `json:"alertRules,omitempty"`
	ApiLinks               []ApiLink      `json:"apiLinks,omitempty"`
	CreatedBy              string         `json:"createdBy,omitempty"`
	CreatedDate            string         `json:"createdDate,omitempty"`
	Description            string         `json:"description,omitempty"`
	Enabled                int            `json:"enabled,omitempty"`
	Groups                 []GroupLabel   `json:"groups,omitempty"`
	LiveShare              int            `json:"liveShare,omitempty"`
	ModifiedBy             string         `json:"modifiedBy,omitempty"`
	ModifiedDate           string         `json:"modifiedDate,omitempty"`
	SavedEvent             int            `json:"savedEvent,omitempty"`
	SharedWithAccounts     []AccountGroup `json:"sharedWithAccounts,omitempty"`
	TestID                 int            `json:"testId,omitempty"`
	TestName               string         `json:"testName,omitempty"`
	Type                   string         `json:"type,omitempty"`
	BgpMeasurements        int            `json:"bgpMeasurements,omitempty"`
	BgpMonitors            []Monitor      `json:"bgpMonitors,omitempty"`
	Interval               int            `json:"interval,omitempty"`
	Direction              string         `json:"direction,omitempty"`
	MtuMeasurements        int            `json:"mtuMeasurements,omitempty"`
	NumPathTraces          int            `json:"numPathTraces,omitempty"`
	Port                   int            `json:"port,omitempty"`
	ProbeMode              string         `json:"probeMode,omitempty"`
	Protocol               string         `json:"protocol,omitempty"`
	Server                 string         `json:"server,omitempty"`
	Dscp                   string         `json:"dscp,omitempty"`
	DscpID                 int            `json:"dscpId,omitempty"`
	Mss                    int            `json:"mss,omitempty"`
	TargetAgentID          int            `json:"targetAgentId,omitempty"`
	ThroughputMeasurements int            `json:"throughputMeasurements,omitempty"`
	ThroughputDuration     int            `json:"throughputDuration,omitempty"`
	ThroughputRate         int            `json:"throughputRate,omitempty"`
}

// AddAgent - Adds an agent to agent test
func (t *AgentAgent) AddAgent(id int) {
	agent := Agent{AgentId: id}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *AgentAgent) AddAlertRule(id int) {
	alertRule := AlertRule{RuleId: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetAgentAgent - Get an agent to agent test
func (c *Client) GetAgentAgent(id int) (*AgentAgent, error) {
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

//DeleteAgentAgent - delete agent to agent test
func (c *Client) DeleteAgentAgent(id int) error {
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
func (c *Client) UpdateAgentAgent(id int, t AgentAgent) (*AgentAgent, error) {
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
