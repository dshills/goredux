package redux

// StateReducer is an interface requiring a Reduce method
type StateReducer interface {
	Reduce(Actioner) StateReducer
	Key() string
}

// Redux is a redux store
type Redux struct {
	subscribers []subscriber
	hooks       []subscriber
	store       *store
}

// New returns a Redux using the State
func New(state ...StateReducer) *Redux {
	r := Redux{store: newStore()}
	for _, s := range state {
		r.store.set(s.Key(), s)
	}
	return &r
}

// State will return the current state
func (r *Redux) State(key string) (StateReducer, bool) {
	return r.store.get(key)
}

// Subscribe will add a listener function to the list of listeners
// when the state changes the function will be called
func (r *Redux) Subscribe(key string, l Listener) {
	r.subscribers = append(r.subscribers, subscriber{listener: l, key: key})
}

// Hook will add a listener for all state changes for all keys
// HookType will describe when to send the state
func (r *Redux) Hook(h Hook, ht HookType) {
	r.hooks = append(r.hooks, subscriber{hook: h, hookType: ht})
}

// Dispatch will run the action against the state and inform the listeners
func (r *Redux) Dispatch(a Actioner) {
	st, ok := r.store.get(a.Key())
	if !ok {
		return
	}

	r.pre(st, a)
	newState := st.Reduce(a)
	r.store.set(a.Key(), newState)
	r.post(newState, a)
	r.notify(a.Key(), newState)
}

func (r *Redux) pre(st StateReducer, a Actioner) {
	for _, sub := range r.hooks {
		if sub.hookType == BeforeReduce {
			sub.hook(st, a)
		}
	}
}

func (r *Redux) post(st StateReducer, a Actioner) {
	for _, sub := range r.hooks {
		if sub.hookType == AfterReduce {
			sub.hook(st, a)
		}
	}
}

func (r *Redux) notify(key string, st StateReducer) {
	for _, sub := range r.subscribers {
		if sub.key == key {
			sub.listener(st)
		}
	}
}
