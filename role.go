package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// AccountGroupRole - an account group role
type AccountGroupRole struct {
	RoleName                 *string      `json:"roleName,omitempty"`
	RoleID                   *int         `json:"roleId,omitempty"`
	HasManagementPermissions *bool        `json:"hasManagementPermissions,omitempty"`
	Builtin                  *bool        `json:"builtin,omitempty"`
	Permissions              []Permission `json:"permissions,omitempty"`
}

// Permission - permission attached to roles
type Permission struct {
	IsManagementPermission *bool   `json:"isManagementPermission"`
	Label                  *string `json:"label"`
	PermissionID           *int    `json:"permissionId"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t AccountGroupRole) MarshalJSON() ([]byte, error) {
	type alias AccountGroupRole

	data, err := json.Marshal((alias)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *AccountGroupRole) UnmarshalJSON(data []byte) error {
	type alias AccountGroupRole
	test := (*alias)(t)

	data, err := jsonIntToBool(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t Permission) MarshalJSON() ([]byte, error) {
	type alias Permission

	data, err := json.Marshal((alias)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *Permission) UnmarshalJSON(data []byte) error {
	type alias Permission
	test := (*alias)(t)

	data, err := jsonIntToBool(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// GetRoles - get roles
func (c *Client) GetRoles() (*[]AccountGroupRole, error) {
	resp, err := c.get("/roles")
	if err != nil {
		return nil, err
	}
	var target map[string][]AccountGroupRole
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	roles := target["roles"]
	return &roles, nil
}

// GetRole - get role
func (c *Client) GetRole(id int) (*AccountGroupRole, error) {
	resp, err := c.get(fmt.Sprintf("/roles/%d", id))
	if err != nil {
		return nil, err
	}
	var target map[string][]AccountGroupRole
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	role := target["roles"][0]
	return &role, nil
}

// DeleteRole - delete role
func (c *Client) DeleteRole(id int) error {
	resp, err := c.post(fmt.Sprintf("/roles/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete role, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateRole - update role
func (c *Client) UpdateRole(id int, role AccountGroupRole) (*AccountGroupRole, error) {
	resp, err := c.post(fmt.Sprintf("/roles/%d/update", id), role, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to update role, response code %d", resp.StatusCode)
	}
	var target AccountGroupRole
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target, nil
}

// CreateRole - create role
func (c *Client) CreateRole(user AccountGroupRole) (*AccountGroupRole, error) {
	resp, err := c.post("/roles/new", user, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to update role, response code %d", resp.StatusCode)
	}
	var target AccountGroupRole
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target, nil
}
