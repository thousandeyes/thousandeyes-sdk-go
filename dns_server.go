package thousandeyes

import (
	"fmt"
)

// Server - a server ?
type Server struct {
	ServerID   int    `json:"serverId,omitempty"`
	ServerName string `json:"serverName,omitempty"`
}

// DNSServer - dns server test
type DNSServer struct {
	Agents                Agents         `json:"agents,omitempty"`
	AlertsEnabled         int            `json:"alertsEnabled,omitempty"`
	AlertRules            AlertRules     `json:"alertRules,omitempty"`
	APILinks              APILinks       `json:"apiLinks,omitempty"`
	CreatedBy             string         `json:"createdBy,omitempty"`
	CreatedDate           string         `json:"createdDate,omitempty"`
	Description           string         `json:"description,omitempty"`
	Enabled               int            `json:"enabled,omitempty"`
	Groups                []GroupLabel   `json:"groups,omitempty"`
	LiveShare             int            `json:"liveShare,omitempty"`
	ModifiedBy            string         `json:"modifiedBy,omitempty"`
	ModifiedDate          string         `json:"modifiedDate,omitempty"`
	SavedEvent            int            `json:"savedEvent,omitempty"`
	SharedWithAccounts    []AccountGroup `json:"sharedWithAccounts,omitempty"`
	TestID                int            `json:"testId,omitempty"`
	TestName              string         `json:"testName,omitempty"`
	Type                  string         `json:"type,omitempty"`
	BandwidthMeasurements int            `json:"bandwidthMeasurements,omitempty"`
	BgpMeasurements       int            `json:"bgpMeasurements,omitempty"`
	BgpMonitors           BGPMonitors    `json:"bgpMonitors,omitempty"`
	DNSServers            []Server       `json:"dnsServers,omitempty"`
	DNSTransportProtocol  string         `json:"dnsTransportProtocol,omitempty"`
	Domain                string         `json:"domain,omitempty"`
	Interval              int            `json:"interval,omitempty"`
	MtuMeasurements       int            `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int            `json:"networkMeasurements,omitempty"`
	NumPathTraces         int            `json:"numPathTraces"`
	PathTraceMode         string         `json:"pathTraceMode,omitempty"`
	ProbeMode             string         `json:"probeMode,omitempty"`
	Protocol              string         `json:"protocol,omitempty"`
	RecursiveQueries      int            `json:"recursiveQueries,omitempty"`
}

// AddAgent - Add dns server test
func (t *DNSServer) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *DNSServer) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
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
