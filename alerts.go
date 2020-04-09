package thousandeyes

import (
	"fmt"
)

// Alerts - list of alerts
type Alerts []Alert

// Alert - An alert
type Alert struct {
	AlertID        int      `json:"alertId,omitempty"`
	TestID         int      `json:"testId,omitempty"`
	TestName       string   `json:"testName,omitempty"`
	Active         int      `json:"active,omitempty"`
	RuleExpression string   `json:"ruleExpression,omitempty"`
	DateStart      string   `json:"dateStart,omitempty"`
	DateEnd        string   `json:"dateEnd,omitempty"`
	ViolationCount int      `json:"violationCount,omitempty"`
	RuleName       string   `json:"ruleName,omitempty"`
	Permalink      string   `json:"permalink,omitempty"`
	Type           string   `json:"type,omitempty"`
	Agents         Agents   `json:"agents,omitempty"`
	Monitors       Monitors `json:"monitors,omitempty"`
	APILinks       APILinks `json:"apiLinks,omitempty"`
}

// AlertRules - list of alert rules
type AlertRules []AlertRule

// AlertRule - An alert rule
type AlertRule struct {
	RuleID                  int         `json:"ruleId"`
	RuleName                string      `json:"ruleName,omitempty"`
	Expression              string      `json:"expression,omitempty"`
	Direction               string      `json:"direction,omitempty"`
	Notifications           interface{} `json:"notifcations,omitempty"`
	NotifyOnClear           int         `json:"notifyOnClear,omitempty"`
	Default                 int         `json:"default,omitempty"`
	AlertType               string      `json:"alertType,omitempty"`
	MinimumSources          int         `json:"minimumSources,omitempty"`
	MinimumSourcesPct       int         `json:"minimumSourcesPct,omitempty"`
	RoundsViolatingOutOf    int         `json:"roundsViolatingOutOf,omitempty"`
	RoundsViolatingRequired int         `json:"roundsViolatingRequired,omitempty"`
	Tests                   int         `json:"tests,omitempty"`
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
	return &target, nil
}

//GetAlertRules - Get alert rules
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

//DeleteAlertRule - delete alert rule
func (c Client) DeleteAlertRule(id int) error {
	resp, err := c.post(fmt.Sprintf("/alert-rules/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete alert rule, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateAlertRule - update alert rule
func (c Client) UpdateAlertRule(id int, a AlertRule) (*AlertRule, error) {
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
