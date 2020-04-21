package thousandeyes

import (
	"fmt"
)

// SIPServer - SIPServer trace test
type SIPServer struct {
	Agents               []Agent        `json:"agents,omitempty"`
	AlertsEnabled        int            `json:"alertsEnabled,omitempty"`
	AlertRules           []AlertRule    `json:"alertRules,omitempty"`
	APILinks             APILinks       `json:"apiLinks,omitempty"`
	CreatedBy            string         `json:"createdBy,omitempty"`
	CreatedDate          string         `json:"createdDate,omitempty"`
	Description          string         `json:"description,omitempty"`
	Enabled              int            `json:"enabled,omitempty"`
	Groups               []GroupLabel   `json:"groups,omitempty"`
	LiveShare            int            `json:"liveShare,omitempty"`
	ModifiedBy           string         `json:"modifiedBy,omitempty"`
	ModifiedDate         string         `json:"modifiedDate,omitempty"`
	SavedEvent           int            `json:"savedEvent,omitempty"`
	SharedWithAccounts   []AccountGroup `json:"sharedWithAccounts,omitempty"`
	TestID               int            `json:"testId,omitempty"`
	TestName             string         `json:"testName,omitempty"`
	Type                 string         `json:"type,omitempty"`
	Interval             int            `json:"interval,omitempty"`
	BgpMeasurements      int            `json:"bgpMeasurements,omitempty"`
	MtuMeasurements      int            `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements  int            `json:"networkMeasurements,omitempty"`
	TargetSipCredentials SIPAuthData    `json:"targetSipCredentials,omitempty"`
	User                 string         `json:"user,omitempty"`
	Password             string         `json:"password,omitempty"`
	OptionsRegex         string         `json:"options_regex,omitempty"`
	RegisterEnabled      int            `json:"registerEnabled,omitempty"`
	SIPTargetTime        int            `json:"sipTargetTime,omitempty"`
	SIPTimeLimit         int            `json:"sipTimeLimit,omitempty"`
}

// AddAgent - Add agemt to sip server  test
func (t *SIPServer) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *SIPServer) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetSIPServer  - get sip server test
func (c *Client) GetSIPServer(id int) (*SIPServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &SIPServer{}, err
	}
	var target map[string][]SIPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreateSIPServer - Create sip server test
func (c Client) CreateSIPServer(t SIPServer) (*SIPServer, error) {
	resp, err := c.post("/tests/sip-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create sip-server test, response code %d", resp.StatusCode)
	}
	var target map[string][]SIPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteSIPServer - delete sip server test
func (c *Client) DeleteSIPServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/sip-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete sip test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateSIPServer - - update sip server test
func (c *Client) UpdateSIPServer(id int, t SIPServer) (*SIPServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/sip-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]SIPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
