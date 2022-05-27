package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// Server - a server ?
type Server struct {
	ServerID   *int    `json:"serverId,omitempty"`
	ServerName *string `json:"serverName,omitempty"`
}

// DNSServer - dns server test
type DNSServer struct {
	// Common test fields
	AlertsEnabled      *bool               `json:"alertsEnabled,omitempty"`
	AlertRules         []AlertRule         `json:"alertRules"`
	APILinks           []APILink           `json:"apiLinks,omitempty"`
	CreatedBy          *string             `json:"createdBy,omitempty"`
	CreatedDate        *string             `json:"createdDate,omitempty"`
	Description        *string             `json:"description,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	Groups             []GroupLabel        `json:"groups,omitempty"`
	ModifiedBy         *string             `json:"modifiedBy,omitempty"`
	ModifiedDate       *string             `json:"modifiedDate,omitempty"`
	SavedEvent         *bool               `json:"savedEvent,omitempty"`
	SharedWithAccounts []SharedWithAccount `json:"sharedWithAccounts,omitempty"`
	TestID             *int64              `json:"testId,omitempty"`
	TestName           *string             `json:"testName,omitempty"`
	Type               *string             `json:"type,omitempty"`
	LiveShare          *bool               `json:"liveShare,omitempty"`

	// Fields unique to this test
	Agents                Agents       `json:"agents,omitempty"`
	BandwidthMeasurements *bool        `json:"bandwidthMeasurements,omitempty"`
	BGPMeasurements       *bool        `json:"bgpMeasurements,omitempty"`
	BGPMonitors           []BGPMonitor `json:"bgpMonitors,omitempty"`
	DNSServers            []Server     `json:"dnsServers,omitempty"`
	DNSTransportProtocol  *string      `json:"dnsTransportProtocol,omitempty"`
	Domain                *string      `json:"domain,omitempty"`
	Interval              *int         `json:"interval,omitempty"`
	MTUMeasurements       *bool        `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   *bool        `json:"networkMeasurements,omitempty"`
	NumPathTraces         *int         `json:"numPathTraces,omitempty"`
	PathTraceMode         *string      `json:"pathTraceMode,omitempty"`
	ProbeMode             *string      `json:"probeMode,omitempty"`
	Protocol              *string      `json:"protocol,omitempty"`
	RecursiveQueries      *bool        `json:"recursiveQueries,omitempty"`
	UsePublicBGP          *bool        `json:"usePublicBgp,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t DNSServer) MarshalJSON() ([]byte, error) {
	type aliasTest DNSServer

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *DNSServer) UnmarshalJSON(data []byte) error {
	type aliasTest DNSServer
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAgent - Add dns server test
func (t *DNSServer) AddAgent(id int) {
	agent := Agent{AgentID: Int(id)}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *DNSServer) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: Int(id)}
	t.AlertRules = append(t.AlertRules, alertRule)
}

//GetDNSServer - get dns server test
func (c *Client) GetDNSServer(id int) (*DNSServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &DNSServer{}, err
	}
	var target map[string][]DNSServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateDNSServer - Create dns server test
func (c Client) CreateDNSServer(t DNSServer) (*DNSServer, error) {
	resp, err := c.post("/tests/dns-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteDNSServer - delete dns server test
func (c *Client) DeleteDNSServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/dns-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete dns server test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateDNSServer - - Update dns server test
func (c *Client) UpdateDNSServer(id int, t DNSServer) (*DNSServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/dns-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
