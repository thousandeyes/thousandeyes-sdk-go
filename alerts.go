package thousandeyes

import (
	"encoding/json"
	"fmt"
	"log"
)

// Alerts - list of alerts
type Alerts []Alert

// Alert - An alert
type Alert struct {
	AlertID        *int64     `json:"alertId,omitempty"`
	TestID         *int64     `json:"testId,omitempty"`
	TestName       *string    `json:"testName,omitempty"`
	Active         *int       `json:"active,omitempty"`
	RuleExpression *string    `json:"ruleExpression,omitempty"`
	DateStart      *string    `json:"dateStart,omitempty"`
	DateEnd        *string    `json:"dateEnd,omitempty"`
	ViolationCount *int       `json:"violationCount,omitempty"`
	RuleName       *string    `json:"ruleName,omitempty"`
	Permalink      *string    `json:"permalink,omitempty"`
	Type           *string    `json:"type,omitempty"`
	Agents         *[]Agent   `json:"agents,omitempty"`
	Monitors       *[]Monitor `json:"monitors,omitempty"`
	APILinks       *[]APILink `json:"apiLinks,omitempty"`
}

// AlertRules - list of alert rules
type AlertRules []AlertRule

// NotificationEmail - Alert Rule Notification Email structure
type NotificationEmail struct {
	Message   *string   `json:"message,omitempty"`
	Recipient *[]string `json:"recipient,omitempty"`
}

// NotificationThirdParty - Alert Rule Notification ThirdParty structure
type NotificationThirdParty struct {
	IntegrationID   *string `json:"integrationId,omitempty"`
	IntegrationName *string `json:"integrationName,omitempty"`
	IntegrationType *string `json:"integrationType,omitempty"`
	Target          *string `json:"target,omitempty"`
	AuthMethod      *string `json:"authMethod,omitempty"`
	AuthUser        *string `json:"authUser,omitempty"`
	AuthToken       *string `json:"authToken,omitempty"`
	Channel         *string `json:"channel,omitempty"`
}

// NotificationWebhook - Alert Rule Notification Webhook structure
type NotificationWebhook struct {
	IntegrationID   *string `json:"integrationId,omitempty"`
	IntegrationName *string `json:"integrationName,omitempty"`
	IntegrationType *string `json:"integrationType,omitempty"`
	Target          *string `json:"target,omitempty"`
}

// Notification - Alert Rule Notification structure
type Notification struct {
	Email      *NotificationEmail        `json:"email,omitempty"`
	ThirdParty *[]NotificationThirdParty `json:"thirdParty,omitempty"`
	Webhook    *[]NotificationWebhook    `json:"webhook,omitempty"`
}

// AlertRule - An alert rule
type AlertRule struct {
	AlertRuleID             *int64         `json:"alertRuleId,omitempty"`
	AlertType               *string        `json:"alertType,omitempty"`
	Default                 *bool          `json:"default,omitempty" te:"int-bool"`
	Direction               *string        `json:"direction,omitempty"`
	Expression              *string        `json:"expression,omitempty"`
	IncludeCoveredPrefixes  *int           `json:"includeCoveredPrefixes,omitempty"`
	MinimumSources          *int           `json:"minimumSources,omitempty"`
	MinimumSourcesPct       *int           `json:"minimumSourcesPct,omitempty"`
	NotifyOnClear           *bool          `json:"notifyOnClear,omitempty" te:"int-bool"`
	RoundsViolatingMode     *string        `json:"roundsViolatingMode,omitempty"`
	RoundsViolatingOutOf    *int           `json:"roundsViolatingOutOf,omitempty"`
	RoundsViolatingRequired *int           `json:"roundsViolatingRequired,omitempty"`
	RuleID                  *int64         `json:"ruleId,omitempty"`
	RuleName                *string        `json:"ruleName,omitempty"`
	Tests                   *[]GenericTest `json:"tests,omitempty"`
	TestIds                 *[]int64       `json:"testIds,omitempty"`
	Notifications           *Notification  `json:"notifications,omitempty"`
	Severity                *string        `json:"severity,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t AlertRule) MarshalJSON() ([]byte, error) {
	type alias AlertRule

	data, err := json.Marshal((alias)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *AlertRule) UnmarshalJSON(data []byte) error {
	type alias AlertRule
	test := (*alias)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// CreateAlertRule - Create alert rule
func (c Client) CreateAlertRule(a AlertRule) (*AlertRule, error) {
	resp, err := c.post("/alert-rules/new", a, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to create alert rule, response code %d", resp.StatusCode)
	}
	var target AlertRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}

	// Set RuleID, because on creation the V6 API returns this as alertRuleId instead.
	// We'll also UNset AlertRuleID so that it isn't seen as a change when it isn't
	// present in other API calls.
	target.RuleID = target.AlertRuleID
	target.AlertRuleID = nil

	return &target, nil
}

// GetAlertRules - Get alert rules
func (c Client) GetAlertRules() (*AlertRules, error) {
	resp, err := c.get(fmt.Sprintf("/alert-rules"))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get alert rule, response code %d", resp.StatusCode)
	}

	var target map[string]AlertRules

	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	alertRules := target["alertRules"]

	return &alertRules, nil
}

// GetAlertRule - Get single alert rule by ID
func (c *Client) GetAlertRule(id int64) (*AlertRule, error) {
	log.Printf("[INFO] Getting Alert Rule %v", id)
	resp, err := c.get(fmt.Sprintf("/alert-rules/%d", id))
	if err != nil {
		return &AlertRule{}, err
	}
	var target map[string][]AlertRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	if len(target["alertRules"]) < 1 {
		return nil, fmt.Errorf("could not get alert rule %v", id)
	}
	return &target["alertRules"][0], nil
}

// DeleteAlertRule - delete alert rule
func (c Client) DeleteAlertRule(id int64) error {
	resp, err := c.post(fmt.Sprintf("/alert-rules/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete alert rule, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateAlertRule - update alert rule
func (c Client) UpdateAlertRule(id int64, a AlertRule) (*AlertRule, error) {
	resp, err := c.post(fmt.Sprintf("/alert-rules/%d/update", id), a, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to update alert rule, response code %d", resp.StatusCode)
	}
	var target AlertRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target, nil
}
