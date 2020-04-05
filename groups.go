package thousandeyes

// AccountGroups - list of account groups
type AccountGroups []AccountGroup

// AccountGroup - an account group
type AccountGroup struct {
	Aid  int
	Name string
}

// GroupLabels - list of group labels
type GroupLabels []GroupLabel

// GroupLabel - group label
type GroupLabel struct {
	GroupName string
	GroupID   int
	BuiltIn   int
}
