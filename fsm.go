package fsm

// State representation of finite state of machine
type State struct {
	current uint32
	prev    uint32
	locked  bool
}

//TransitionTo method set the new state(t) of the machine
func (s State) TransitionTo(newState uint32) uint32 {
	if s.locked == false {
		s.prev = s.current
		s.current = newState
	}
	return s.current
}

// Lock method the state and transitions can't happen
func (s State) Lock() {
	s.locked = true
}

// Unlock method the state and transitions can happen
func (s State) Unlock() {
	s.locked = false
}
