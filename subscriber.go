package redux

// Listener is a function that receives State after it changes
type Listener func(StateReducer)

// Hook is a function that receives State and an Action
type Hook func(StateReducer, Actioner)

// HookType is a constant describing when to perform the hook function
type HookType int

// HookType constants
const (
	BeforeReduce HookType = iota
	AfterReduce
)

type subscriber struct {
	listener Listener
	key      string
	hook     Hook
	hookType HookType
}
