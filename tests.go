package thousandeyes

import "fmt"

type GenericTest struct {
	Agents             []Agent        `json:"agents,omitempty"`
	AlertsEnabled      int            `json:"alertsEnabled,omitempty"`
	AlertRules         []AlertRule    `json:"alertRules,omitempty"`
	APILinks           APILinks       `json:"apiLinks,omitempty"`
	CreatedBy          string         `json:"createdBy,omitempty"`
	CreatedDate        string         `json:"createdDate,omitempty"`
	Description        string         `json:"description,omitempty"`
	Enabled            int            `json:"enabled,omitempty"`
	Groups             []GroupLabel   `json:"groups,omitempty"`
	ModifiedBy         string         `json:"modifiedBy,omitempty"`
	ModifiedDate       string         `json:"modifiedDate,omitempty"`
	SavedEvent         int            `json:"savedEvent,omitempty"`
	SharedWithAccounts []AccountGroup `json:"sharedWithAccounts,omitempty"`
	TestID             int            `json:"testId,omitempty"`
	TestName           string         `json:"testName,omitempty"`
	Type               string         `json:"type,omitempty"`
}

// GetTests  - get all tests
func (c *Client) GetTests() (*[]GenericTest, error) {
	resp, err := c.get("/tests")
	if err != nil {
		return nil, err
	}
	var target map[string][]GenericTest
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	tests := target["test"]
	return &tests, nil
}

// GetTest - Get test
func (c *Client) GetTest(id int) (*GenericTest, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return nil, err
	}
	var target map[string][]GenericTest
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	test := target["test"][0]
	return &test, nil
}
