package thousandeyes

// AccountGroups - list of account groups
type AccountGroups []AccountGroup

// AccountGroup - An account within a ThousandEyes organization
type AccountGroup struct {
	AccountGroupName string `json:"accountGroupName,omitempty"`
	AID              int    `json:"aid,omitempty"`
}

// SharedWithAccount describes accounts with which a resource is shared.
// This is separate from the AccountGroup above only due to the difference
// in JSON object names.
type SharedWithAccount struct {
	AccountGroupName string `json:"name,omitempty"`
	AID              int    `json:"aid,omitempty"`
}
