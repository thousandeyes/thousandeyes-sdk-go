package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// FTPServer - ftp server test
type FTPServer struct {
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
	Agents              *[]Agent `json:"agents,omitempty"`
	BGPMeasurements     *bool    `json:"bgpMeasurements,omitempty" te:"int-bool"`
	DownloadLimit       *int     `json:"downloadLimit,omitempty"`
	FTPTargetTime       *int     `json:"ftpTargetTime,omitempty"`
	FTPTimeLimit        *int     `json:"ftpTimeLimit,omitempty"`
	Interval            *int     `json:"interval,omitempty"`
	MTUMeasurements     *bool    `json:"mtuMeasurements,omitempty" te:"int-bool"`
	NetworkMeasurements *bool    `json:"networkMeasurements,omitempty" te:"int-bool"`
	NumPathTraces       *int     `json:"numPathTraces,omitempty"`
	Password            *string  `json:"password,omitempty"`
	PathTraceMode       *string  `json:"pathTraceMode,omitempty"`
	ProbeMode           *string  `json:"probeMode,omitempty"`
	Protocol            *string  `json:"protocol,omitempty"`
	RequestType         *string  `json:"requestType,omitempty"`
	URL                 *string  `json:"url,omitempty"`
	UseActiveFTP        *int     `json:"useActiveFtp,omitempty"`
	UseExplicitFTPS     *int     `json:"useExplicitFtps,omitempty"`
	Username            *string  `json:"username,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t FTPServer) MarshalJSON() ([]byte, error) {
	type aliasTest FTPServer

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *FTPServer) UnmarshalJSON(data []byte) error {
	type aliasTest FTPServer
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAgent - Add ftp server test
func (t *FTPServer) AddAgent(id int) {
	agent := Agent{AgentID: Int(id)}
	*t.Agents = append(*t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *FTPServer) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: Int(id)}
	*t.AlertRules = append(*t.AlertRules, alertRule)
}

// GetFTPServer - get ftp server test
func (c *Client) GetFTPServer(id int) (*FTPServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &FTPServer{}, err
	}
	var target map[string][]FTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateFTPServer - Create ftp server test
func (c Client) CreateFTPServer(t FTPServer) (*FTPServer, error) {
	resp, err := c.post("/tests/ftp-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create ftp test, response code %d", resp.StatusCode)
	}
	var target map[string][]FTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// DeleteFTPServer - delete ftp server test
func (c *Client) DeleteFTPServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/ftp-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete ftp server test, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateFTPServer - - Update ftp server test
func (c *Client) UpdateFTPServer(id int, t FTPServer) (*FTPServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/ftp-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update ftp test, response code %d", resp.StatusCode)
	}
	var target map[string][]FTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
