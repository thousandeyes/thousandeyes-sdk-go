package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// GenericTest - GenericTest struct to represent all test types
type GenericTest struct {
	// Common test fields
	AlertsEnabled      *bool               `json:"alertsEnabled,omitempty" te:"int-bool"`
	AlertRules         []AlertRule         `json:"alertRules"`
	APILinks           []APILink           `json:"apiLinks,omitempty"`
	CreatedBy          *string             `json:"createdBy,omitempty"`
	CreatedDate        *string             `json:"createdDate,omitempty"`
	Description        *string             `json:"description,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty" te:"int-bool"`
	Groups             []GroupLabel        `json:"groups,omitempty"`
	ModifiedBy         *string             `json:"modifiedBy,omitempty"`
	ModifiedDate       *string             `json:"modifiedDate,omitempty"`
	SavedEvent         *bool               `json:"savedEvent,omitempty" te:"int-bool"`
	SharedWithAccounts []SharedWithAccount `json:"sharedWithAccounts,omitempty"`
	TestID             *int64              `json:"testId,omitempty"`
	TestName           *string             `json:"testName,omitempty"`
	Type               *string             `json:"type,omitempty"`
	// Fields unique to this test
	Agents []Agent `json:"agents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t GenericTest) MarshalJSON() ([]byte, error) {
	type aliasTest GenericTest

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *GenericTest) UnmarshalJSON(data []byte) error {
	type aliasTest GenericTest
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// GetTests  - get all tests
func (c *Client) GetTests() (*[]GenericTest, error) {
	resp, err := c.get("/tests")
	if err != nil {
		return nil, err
	}
	var target map[string][]GenericTest
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
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
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	test := target["test"][0]
	return &test, nil
}
