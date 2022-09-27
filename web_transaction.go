package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// WebTransaction - a web transcation test
type WebTransaction struct {
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
	Agents                *[]Agent       `json:"agents,omitempty"`
	AuthType              *string        `json:"authType,omitempty"`
	BandwidthMeasurements *bool          `json:"bandwidthMeasurements,omitempty" te:"int-bool"`
	ContentRegex          *string        `json:"contentRegex,omitempty"`
	Credentials           *[]int         `json:"credentials,omitempty"`
	CustomHeaders         *CustomHeaders `json:"customHeaders,omitempty"`
	DesiredStatusCode     *string        `json:"desiredStatusCode,omitempty"`
	HTTPTargetTime        *int           `json:"httpTargetTime,omitempty"`
	HTTPTimeLimit         *int           `json:"httpTimeLimit,omitempty"`
	HTTPVersion           *int           `json:"httpVersion,omitempty"`
	IncludeHeaders        *bool          `json:"includeHeaders,omitempty" te:"int-bool"`
	Interval              *int           `json:"interval,omitempty"`
	MTUMeasurements       *bool          `json:"mtuMeasurements,omitempty" te:"int-bool"`
	NetworkMeasurements   *bool          `json:"networkMeasurements,omitempty" te:"int-bool"`
	NumPathTraces         *int           `json:"numPathTraces,omitempty"`
	Password              *string        `json:"password,omitempty"`
	PathTraceMode         *string        `json:"pathTraceMode,omitempty"`
	ProbeMode             *string        `json:"probeMode,omitempty"`
	Protocol              *string        `json:"protocol,omitempty"`
	SSLVersionID          *int64         `json:"sslVersionId,omitempty"`
	SubInterval           *int           `json:"subinterval,omitempty"`
	TargetTime            *int           `json:"targetTime,omitempty"`
	TimeLimit             *int           `json:"timeLimit,omitempty"`
	TransactionScript     *string        `json:"transactionScript,omitempty"`
	URL                   *string        `json:"url,omitempty"`
	UseNTLM               *bool          `json:"useNtlm,omitempty" te:"int-bool"`
	UserAgent             *string        `json:"userAgent,omitempty"`
	Username              *string        `json:"username,omitempty"`
	VerifyCertificate     *bool          `json:"verifyCertificate,omitempty" te:"int-bool"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t WebTransaction) MarshalJSON() ([]byte, error) {
	type aliasTest WebTransaction

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *WebTransaction) UnmarshalJSON(data []byte) error {
	type aliasTest WebTransaction
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// CreateWebTransaction - Create a web transaction test
func (c Client) CreateWebTransaction(t WebTransaction) (*WebTransaction, error) {
	resp, err := c.post("/tests/web-transactions/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create web transaction, response code %d", resp.StatusCode)
	}
	var target map[string][]WebTransaction
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// GetWebTransaction - get a web transactiont test
func (c *Client) GetWebTransaction(id int64) (*WebTransaction, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &WebTransaction{}, err
	}
	var target map[string][]WebTransaction
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// DeleteWebTransaction - delete a web transactiont est
func (c *Client) DeleteWebTransaction(id int64) error {
	resp, err := c.post(fmt.Sprintf("/tests/web-transactions/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete http server, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateWebTransaction - update a web transaction test
func (c *Client) UpdateWebTransaction(id int64, t WebTransaction) (*WebTransaction, error) {
	resp, err := c.post(fmt.Sprintf("/tests/web-transactions/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to web transaction, response code %d", resp.StatusCode)
	}
	var target map[string][]WebTransaction
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
