package goredux

// WrappedReducerFunc is a function that gets the state and an action
// and returns a new state
type WrappedReducerFunc func(interface{}, Actioner) interface{}

// WrappedState is a wrapper for an arbitrary interface{}
// that satisifies the StateReducer interface
type WrappedState struct {
	State      interface{}
	KeyName    string
	ReduceFunc WrappedReducerFunc
}

// WrapState will return a StateReducer for an interface{} representing state
func WrapState(state interface{}, key string, fx WrappedReducerFunc) *WrappedState {
	return &WrappedState{
		State:      state,
		KeyName:    key,
		ReduceFunc: fx,
	}
}

// Key returns the store key for the state
func (w *WrappedState) Key() string {
	return w.KeyName
}

// Reduce takes an action and returns a new state
func (w *WrappedState) Reduce(a Actioner) StateReducer {
	return &WrappedState{
		State:      w.ReduceFunc(w.State, a),
		KeyName:    w.KeyName,
		ReduceFunc: w.ReduceFunc,
	}
}
