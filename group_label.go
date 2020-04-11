package thousandeyes

import "fmt"

// GroupLabels - list of labels
type GroupLabels []GroupLabel

// GroupLabel - label
type GroupLabel struct {
	GroupLabelName string      `json:"name,omitempty"`
	GroupLabelID   int         `json:"groupId,omitempty"`
	BuiltIn        int         `json:"builtin,omitempty"`
	GroupLabelType string      `json:"type,omitempty"`
	Agents         Agents      `json:"agents,omitempty"`
	Tests          interface{} `json:"tests,omitempty"`
}

// GetGroupLabels - Get labels
func (c *Client) GetGroupLabels() (*GroupLabels, error) {
	resp, err := c.get("/groups")
	if err != nil {
		return nil, err
	}
	var target map[string]GroupLabels
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	labels := target["groups"]
	return &labels, nil
}

// GetGroupLabelsByType - Get label by type
func (c *Client) GetGroupLabelsByType(t string) (*GroupLabels, error) {
	resp, err := c.get("/groups/" + t)
	if err != nil {
		return &GroupLabels{}, err
	}
	var target map[string]GroupLabels
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	labels := target["groups"]
	return &labels, nil
}

// GetGroupLabelsByID - Get label
func (c *Client) GetGroupLabelsByID(id int) (*GroupLabels, error) {
	resp, err := c.get(fmt.Sprintf("/groups/%d", id))
	if err != nil {
		return &GroupLabels{}, err
	}
	var target map[string]GroupLabels
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	labels := target["groups"]
	return &labels, nil
}

// CreateGroupLabel - Create label
func (c Client) CreateGroupLabel(a GroupLabel) (*GroupLabels, error) {
	resp, err := c.post("/groups/new", a, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to create label, response code %d", resp.StatusCode)
	}

	var target map[string]GroupLabels

	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	labels := target["groups"]

	return &labels, nil
}

//DeleteGroupLabel - delete label
func (c Client) DeleteGroupLabel(id int) error {
	resp, err := c.post(fmt.Sprintf("/groups/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete label, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateGroupLabel - update label
func (c Client) UpdateGroupLabel(id int, a GroupLabel) (*GroupLabels, error) {
	resp, err := c.post(fmt.Sprintf("/groups/%d/update", id), a, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to update label, response code %d", resp.StatusCode)
	}

	var target map[string]GroupLabels
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	labels := target["groups"]
	return &labels, nil

}
