package redux

import "sync"

type store struct {
	state map[string]StateReducer
	mutex *sync.Mutex
}

func newStore() *store {
	return &store{
		state: make(map[string]StateReducer),
		mutex: &sync.Mutex{},
	}
}

func (s *store) get(key string) (StateReducer, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	v, ok := s.state[key]
	return v, ok
}

func (s *store) set(key string, i StateReducer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.state[key] = i
}
