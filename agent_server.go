package thousandeyes

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// AgentServer  - Agent to server test
type AgentServer struct {
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
	Agents                *[]Agent      `json:"agents,omitempty"`
	BandwidthMeasurements *bool         `json:"bandwidthMeasurements,omitempty" te:"int-bool"`
	BGPMeasurements       *bool         `json:"bgpMeasurements,omitempty" te:"int-bool"`
	BGPMonitors           *[]BGPMonitor `json:"bgpMonitors,omitempty"`
	Interval              *int          `json:"interval,omitempty"`
	MTUMeasurements       *bool         `json:"mtuMeasurements,omitempty" te:"int-bool"`
	NetworkMeasurements   *bool         `json:"networkMeasurements,omitempty" te:"int-bool"`
	NumPathTraces         *int          `json:"numPathTraces,omitempty"`
	PathTraceMode         *string       `json:"pathTraceMode,omitempty"`
	Port                  *int          `json:"port,omitempty"`
	ProbeMode             *string       `json:"probeMode,omitempty"`
	Protocol              *string       `json:"protocol,omitempty"`
	Server                *string       `json:"server,omitempty"`
	UsePublicBGP          *bool         `json:"usePublicBgp,omitempty" te:"int-bool"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t AgentServer) MarshalJSON() ([]byte, error) {
	type aliasTest AgentServer

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *AgentServer) UnmarshalJSON(data []byte) error {
	type aliasTest AgentServer
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// extractPort - Set Server and Port fields if they are combined in the Server field.
func extractPort(test AgentServer) (AgentServer, error) {
	// Unfortunately, the V6 API returns the server value with the port,
	// rather than having them in separate values as the API requires for
	// submissions.  Not required for ICMP tests.
	if (test.Protocol == nil || *test.Protocol != "ICMP") &&
		(test.Server != nil && strings.Index(*test.Server, ":") != -1) {

		serverParts := strings.Split(*test.Server, ":")
		*test.Server = serverParts[0]
		port, err := strconv.Atoi(serverParts[1])
		if err != nil {
			err = fmt.Errorf("Invalid port in server declaration")
			return test, err
		}
		test.Port = Int(port)
	}
	return test, nil
}

// AddAgent - Add agent to server test
func (t *AgentServer) AddAgent(id int64) {
	agent := Agent{AgentID: Int64(id)}
	*t.Agents = append(*t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *AgentServer) AddAlertRule(id int64) {
	alertRule := AlertRule{RuleID: Int64(id)}
	*t.AlertRules = append(*t.AlertRules, alertRule)
}

// GetAgentServer - Get agent to server test
func (c *Client) GetAgentServer(id int64) (*AgentServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &AgentServer{}, err
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	test := target["test"][0]
	test, err = extractPort(test)
	if err != nil {
		return nil, err
	}
	return &test, nil
}

// CreateAgentServer  - Create agent to server test
func (c Client) CreateAgentServer(t AgentServer) (*AgentServer, error) {
	resp, err := c.post("/tests/agent-to-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create agent server, response code %d", resp.StatusCode)
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	test := target["test"][0]
	test, err = extractPort(test)
	if err != nil {
		return nil, err
	}
	return &test, nil
}

// DeleteAgentServer  - Delete agent to server test
func (c *Client) DeleteAgentServer(id int64) error {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete agent server, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateAgentServer  - Update agent to server test
func (c *Client) UpdateAgentServer(id int64, t AgentServer) (*AgentServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update agent server, response code %d", resp.StatusCode)
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
