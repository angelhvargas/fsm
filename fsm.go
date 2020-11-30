package fsm

// State representation of finite state of machine
type State struct {
	current uint32
	prev    uint32
}

//TransitionTo method set the new state(t) of the machine
func (s State) TransitionTo(newState uint32) uint32 {
	s.prev = s.current
	s.current = newState
	return s.current
}
