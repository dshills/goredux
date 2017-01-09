package redux

// Actioner is a Redux Action
type Actioner interface {
	Action() string
	Key() string
}

// BasicAction is an action that only contains an actions string
// if no other data is required in the action this is the action to use
type BasicAction struct {
	Act     string
	KeyName string
}

// Action returns the action string
func (a *BasicAction) Action() string {
	return a.Act
}

// Key returns the action key
func (a *BasicAction) Key() string {
	return a.KeyName
}
