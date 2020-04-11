package thousandeyes

// AccountGroups - list of account groups
type AccountGroups []AccountGroup

// AccountGroup - an account group
type AccountGroup struct {
	Aid  int    `json:"aid"`
	Name string `json:"name"`
}
