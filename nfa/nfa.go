package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise
	if input == nil {
		if start != final {
			return false
		}
		return true
	}
	s_temp := transitions(start, input[0])
	var s []state
	for i:= 0; i < len(s_temp); i++ {
		s = append(s, s_temp[i])
	}
	for i:= 1; i < len(input); i++ {
		var t []state
		for j:= 0; j < len(s); j++ {
			t_temp := transitions(s[j],input[i])
			for k:= 0; k < len(t_temp); k++ {
				t = append(t, t_temp[k])
			}
		}
		s = t
	}
	for j:=0; j < len(s); j++ {
		if s[j] == final {
			return true
		}
	}
	return false
}
