package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// SIPAuthData - Authentication fields for SIP tests
type SIPAuthData struct {
	AuthUser     *string `json:"authUser,omitempty"`
	Password     *string `json:"password,omitempty"`
	Port         *int    `json:"port,omitempty"`
	Protocol     *string `json:"protocol,omitempty"`
	SIPProxy     *string `json:"sipProxy,omitempty"`
	SIPRegistrar *string `json:"sipRegistrar,omitempty"`
	User         *string `json:"user,omitempty"`
}

// VoiceCall - VoiceCall trace test
type VoiceCall struct {
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
	Agents               *[]Agent     `json:"agents,omitempty"`
	BGPMeasurements      *bool        `json:"bgpMeasurements,omitempty" te:"int-bool"`
	Codec                *string      `json:"codec,omitempty"`
	CodecID              *int         `json:"codecId,omitempty"`
	DSCP                 *string      `json:"dscp,omitempty"`
	DSCPID               *int         `json:"dscpId,omitempty"`
	Duration             *int         `json:"duration,omitempty"`
	Interval             *int         `json:"interval,omitempty"`
	JitterBuffer         *int         `json:"jitterBuffer,omitempty"`
	NumPathTraces        *int         `json:"numPathTraces,omitempty"`
	SIPTargetTime        *int         `json:"sipTargetTime,omitempty"`
	SIPTimeLimit         *int         `json:"sipTimeLimit,omitempty"`
	SourceSIPCredentials *SIPAuthData `json:"sourceSipCredentials,omitempty"`
	TargetAgentID        *int         `json:"targetAgentId,omitempty"`
	TargetSIPCredentials *SIPAuthData `json:"targetSipCredentials,omitempty"`
	UsePublicBGP         *bool        `json:"usePublicBgp,omitempty" te:"int-bool"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t VoiceCall) MarshalJSON() ([]byte, error) {
	type aliasTest VoiceCall

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *VoiceCall) UnmarshalJSON(data []byte) error {
	type aliasTest VoiceCall
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAgent - Add agent to voice call  test
func (t *VoiceCall) AddAgent(id int) {
	agent := Agent{AgentID: Int(id)}
	*t.Agents = append(*t.Agents, agent)
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
		return &t, fmt.Errorf("failed to create voice-call test, response code %d", resp.StatusCode)
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
