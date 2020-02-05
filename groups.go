package thousandeyes

type AccountGroups []AccountGroup

type AccountGroup struct {
	Aid  int
	Name string
}

type GroupLabels []GroupLabel

type GroupLabel struct {
	GroupName string
	GroupId   int
	BuiltIn   int
}
