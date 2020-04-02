package thousandeyes

import (
	"fmt"
)

type Alert struct {
	AlertId        int       `json:"alertId,omitempty"`
	TestId         int       `json:"testId,omitempty"`
	TestName       string    `json:"testName,omitempty"`
	Active         int       `json:"active,omitempty"`
	RuleExpression string    `json:"ruleExpression,omitempty"`
	DateStart      string    `json:"dateStart,omitempty"`
	DateEnd        string    `json:"dateEnd,omitempty"`
	ViolationCount int       `json:"violationCount,omitempty"`
	RuleName       string    `json:"ruleName,omitempty"`
	Permalink      string    `json:"permalink,omitempty"`
	Type           string    `json:"type,omitempty"`
	Agents         []Agent   `json:"agents,omitempty"`
	Monitors       []Monitor `json:"monitors,omitempty"`
	ApiLinks       []ApiLink `json:"apiLinks,omitempty"`
}

type AlertRule struct {
	RuleId                  int         `json:"ruleId,omitempty"`
	RuleName                string      `json:"ruleName,omitempty"`
	Expression              string      `json:"expression,omitempty"`
	Direction               string      `json:"direction,omitempty"`
	Notifications           interface{} `json:"notifcations,omitempty"`
	NotifyOnClear           bool        `json:"notifyOnClear,omitempty"`
	Default                 bool        `json:"default,omitempty"`
	AlertType               string      `json:"alertType,omitempty"`
	MinimumSources          int         `json:"minimumSources,omitempty"`
	MinimumSourcesPct       int         `json:"minimumSourcesPct,omitempty"`
	RoundsViolatingOutOf    int         `json:"roundsViolatingOutOf,omitempty"`
	RoundsViolatingRequired int         `json:"roundsViolatingRequired,omitempty"`
	Tests                   int         `json:"tests,omitempty"`
}

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

func (c Client) GetAlertRule(id int) (*AlertRule, error) {
	resp, err := c.get(fmt.Sprintf("/alert-rules/%d", id))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get alert rule, response code %d", resp.StatusCode)
	}
	var target AlertRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target, nil
}

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
