package fsm

import (
	"testing"
)

func TestState(t *testing.T) {
	// test cases for State transitions
	var tests = []struct {
		startState    uint32
		toState       uint32
		expectedState uint32
	}{
		{1, 1, 1},
		{1, 2, 2},
		{2, 1, 1},
	}

	for _, test := range tests {
		s := State{test.startState, test.startState, false}
		if res := s.TransitionTo(test.startState); res != s.current {
			t.Error("error")
		}
	}

}
