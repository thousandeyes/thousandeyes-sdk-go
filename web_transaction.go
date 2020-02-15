package thousandeyes

import (
	"errors"
	"fmt"
)

type WebTransaction struct {
	Agents                Agents              `json:"agents,omitempty"`
	AlertsEnabled         int                 `json:"alertsEnabled,omitempty"`
	AlertRules            []AlertRule         `json:"alertRules,omitempty"`
	ApiLinks              []ApiLink           `json:"apiLinks,omitempty"`
	CreatedBy             string              `json:"createdBy,omitempty"`
	CreatedDate           string              `json:"createdDate,omitempty"`
	Description           string              `json:"description,omitempty"`
	Enabled               int                 `json:"enabled,omitempty"`
	Groups                []GroupLabels       `json:"groups,omitempty"`
	LiveShare             int                 `json:"liveShare,omitempty"`
	ModifiedBy            string              `json:"modifiedBy,omitempty"`
	ModifiedDate          string              `json:"modifiedDate,omitempty"`
	SavedEvent            int                 `json:"savedEvent,omitempty"`
	SharedWithAccounts    []AccountGroup      `json:"sharedWithAccounts,omitempty"`
	TestId                int                 `json:"testId,omitempty"`
	TestName              string              `json:"testName,omitempty"`
	Type                  string              `json:"type,omitempty"`
	AuthType              string              `json:"authType,omitempty"`
	BandwidthMeasurements int                 `json:"bandwidthMeasurements,omitempty"`
	ContentRegex          string              `json:"contentRegex,omitempty"`
	Credentials           []int               `json:"credentials,omitempty"`
	CustomHeaders         []map[string]string `json:"customHeaders,omitempty"`
	DesiredStatusCode     string              `json:"desiredStatusCode,omitempty"`
	HttpTargetTime        int                 `json:"httpTargetTime,omitempty"`
	HttpTimeLimit         int                 `json:"httpTimeLimit,omitempty"`
	HttpVersion           int                 `json:"httpVersion,omitempty"`
	IncludeHeaders        int                 `json:"ncludeHeaders,omitempty"`
	Interval              int                 `json:"interval,omitempty"`
	MtuMeasurements       int                 `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int                 `json:"networkMeasurements,omitempty"`
	NumPathTraces         int                 `json:"numPathTraces,omitempty"`
	Password              string              `json:"password,omitempty"`
	ProbeMode             string              `json:"probeMode,omitempty"`
	Protocol              string              `json:"protocol,omitempty"`
	SslVersionId          int                 `json:"sslVersionId,omitempty"`
	Subinterval           int                 `json:"subInterval,omitempty"`
	TargetTime            int                 `json:"targetTime,omitempty"`
	TimeLimit             int                 `json:"timeLimit,omitempty"`
	TransactionScript     string              `json:"transactionScript,omitempty"`
	Url                   string              `json:"url,omitempty"`
	UseNtlm               int                 `json:"useNtlm,omitempty"`
	UserAgent             string              `json:"userAgent,omitempty"`
	Username              string              `json:"username,omitempty"`
	VerifyCertificate     int                 `json:"verifyCertificate,omitempty"`
}

func (c Client) CreateWebTransaction(t WebTransaction) (*WebTransaction, error) {
	resp, err := c.post("/tests/web-transactions/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, errors.New(fmt.Sprintf("failed to create web transaction, response code %d", resp.StatusCode))
	}
	var target map[string][]WebTransaction
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

func (c *Client) GetWebTransaction(id int) (*WebTransaction, error) {
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

func (c *Client) DeleteWebTransaction(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/web-transactions/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return errors.New(fmt.Sprintf("failed to delete http server, response code %d", resp.StatusCode))
	}
	return nil
}

func (c *Client) UpdateWebTransaction(id int, t WebTransaction) (*WebTransaction, error) {
	resp, err := c.post(fmt.Sprintf("/tests/web-transactions/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, errors.New(fmt.Sprintf("failed to web transaction, response code %d", resp.StatusCode))
	}
	var target map[string][]WebTransaction
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
