package thousandeyes

import (
	"fmt"
)

type sipAuthData struct {
	AuthUser     string `json:"authUser,omitempty"`
	Password     string `json:"password,omitempty"`
	Port         int    `json:"port,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	SipProxy     string `json:"sipProxy,omitempty"`
	SipRegistrar string `json:"sipRegistrar,omitempty"`
	User         string `json:"user,omitempty"`
}

// VoiceCall - VoiceCall trace test
type VoiceCall struct {
	Agents               []Agent        `json:"agents,omitempty"`
	AlertsEnabled        int            `json:"alertsEnabled,omitempty"`
	AlertRules           []AlertRule    `json:"alertRules,omitempty"`
	APILinks             []ApiLink      `json:"apiLinks,omitempty"`
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
	Dscp                 string         `json:"dscp,omitempty"`
	DscpID               int            `json:"dscpId,omitempty"`
	TargetSipCredentials sipAuthData    `json:"targetSipCredentials,omitempty"`
	SourceSipCredentials sipAuthData    `json:"sourceSipCredentials,omitempty"`
	Codec                string         `json:"codec,omitempty"`
	CodecID              int            `json:"codecId,omitempty"`
	Duration             int            `json:"duration,omitempty"`
	JitterBuffer         int            `json:"jitterBuffer,omitempty"`
	SIPTargetTime        int            `json:"sipTargetTime,omitempty"`
	SIPTimeLimit         int            `json:"sipTimeLimit,omitempty"`
	TargetAgentID        int            `json:"targetAgentId,omitempty"`
}

// AddAgent - Add agent to voice call  test
func (t *VoiceCall) AddAgent(id int) {
	agent := Agent{AgentId: id}
	t.Agents = append(t.Agents, agent)
}

// GetVoiceCall  - get voice call test
func (c *Client) GetVoiceCall(id int) (*VoiceCall, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &VoiceCall{}, err
	}
	var target map[string][]VoiceCall
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreateVoiceCall - Create voice call test
func (c Client) CreateVoiceCall(t VoiceCall) (*VoiceCall, error) {
	resp, err := c.post("/tests/voice-call/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create sip-server test, response code %d", resp.StatusCode)
	}
	var target map[string][]VoiceCall
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteVoiceCall - delete voice call test
func (c *Client) DeleteVoiceCall(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/voice-call/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete voice-call test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateVoiceCall - - update voice call test
func (c *Client) UpdateVoiceCall(id int, t VoiceCall) (*VoiceCall, error) {
	resp, err := c.post(fmt.Sprintf("/tests/voice-call/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]VoiceCall
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
