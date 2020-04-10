package thousandeyes

import "fmt"

// Labels - list of labels
type Labels []Label

// Label - label
type Label struct {
	LabelName string `json:"name,omitempty"`
	LabelID   int    `json:"groupId,omitempty"`
	BuiltIn   int    `json:"builtin,omitempty"`
	Type      string `json:"type,omitempty"`
	Agents    Agents `json:"agents,omitempty"`
	// Tests     Tests  `json:"tests,omitempty"`  # https://github.com/william20111/go-thousandeyes/issues/49
}

// GetLabels - Get labels
func (c *Client) GetLabels() (*Labels, error) {
	resp, err := c.get("/groups")
	if err != nil {
		return &Labels{}, err
	}
	var target map[string]Labels
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	labels := target["groups"]
	return &labels, nil
}

// GetLabelByType - Get label by type
func (c *Client) GetLabelByType(t string) (*Labels, error) {
	resp, err := c.get("/groups/" + t)
	if err != nil {
		return &Labels{}, err
	}
	var target map[string]Labels
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	labels := target["groups"]
	return &labels, nil
}

// GetLabel - Get label
func (c *Client) GetLabel(id int) (*Labels, error) {
	resp, err := c.get("/groups/" + string(id))
	if err != nil {
		return &Labels{}, err
	}
	var target map[string]Labels
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	labels := target["groups"]
	return &labels, nil
}

// CreateLabel - Create label
func (c Client) CreateLabel(a Label) (*Label, error) {
	resp, err := c.post("/groups/new", a, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to create label, response code %d", resp.StatusCode)
	}
	var target Label
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target, nil
}

//DeleteLabel - delete label
func (c Client) DeleteLabel(id int) error {
	resp, err := c.post(fmt.Sprintf("/groups/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete label, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateLabel - update label
func (c Client) UpdateLabel(id int, a Label) (*Label, error) {
	resp, err := c.post(fmt.Sprintf("/groups/%d/update", id), a, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to update label, response code %d", resp.StatusCode)
	}
	var target Label
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target, nil
}
