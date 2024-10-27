package main

import (
	"fmt"
	"sort"
)

// FiniteState struct
type FiniteState struct {
	nextState      int
	CurrentState   int
	TerminalStates []int
	Transitions    map[int]map[rune]int
}

func NewAutomata() *FiniteState {
	return &FiniteState{
		nextState:      1,
		CurrentState:   0,
		TerminalStates: []int{0},
		Transitions:    make(map[int]map[rune]int),
	}
}

func Create(chars []rune) *FiniteState {
	f := NewAutomata()
	f.AddTransition(0, 1, chars)
	f.TerminalStates = []int{1}
	return f
}

func (f *FiniteState) GetTerminalStates() []int {
	return f.TerminalStates
}

func (f *FiniteState) addTerminal(terminal int) {
	for _, val := range f.TerminalStates {
		if val == terminal {
			return
		}
	}
	f.TerminalStates = append(f.TerminalStates, terminal)
}

func (f *FiniteState) AddTransition(from, to int, chars []rune) {

	//Update the next state indicator if necessary
	if from > f.nextState {
		f.nextState = from + 1
	}

	if to > f.nextState {
		f.nextState = to + 1
	}

	// If we have a transition set from this state already
	// then add/update
	if transitionsFrom, ok := f.Transitions[from]; ok {
		for _, ch := range chars {
			if f.isTerminal(transitionsFrom[ch]) {
				f.addTerminal(to)
			}
			transitionsFrom[ch] = to
		}
	} else {
		transitionsFrom := make(map[rune]int)
		for _, ch := range chars {
			transitionsFrom[ch] = to
		}
		f.Transitions[from] = transitionsFrom
	}
}

// Append the given automata onto the end of this one
func (f *FiniteState) Append(other *FiniteState) {
	offset := f.nextState
	f.nextState += other.nextState

	//Update transitions from the other initial
	for from, transition := range other.Transitions {
		if from == 0 {
			for _, terminal := range f.TerminalStates {
				for ch, to := range transition {
					f.AddTransition(terminal, to+offset, []rune{ch})
				}
			}
		} else {
			for ch, to := range transition {
				if to == 0 {
					for terminal := range f.TerminalStates {
						f.AddTransition(from+offset, terminal, []rune{ch})
					}
				} else {
					f.AddTransition(from+offset, to+offset, []rune{ch})
				}
			}
		}
	}

	//Set new terminal
	newTerms := []int{}
	for _, term := range other.TerminalStates {
		if term == 0 {
			for _, t := range f.TerminalStates {
				newTerms = append(newTerms, t)
			}
		}
		newTerms = append(newTerms, term+offset)
	}
	f.TerminalStates = newTerms

}

// Union the given automata with this one
func (f *FiniteState) Union(other *FiniteState) {
	offset := f.nextState
	f.nextState += other.nextState

	mapping := make(map[int]int)

	//Copy transitions from other
	for from, transition := range other.Transitions {

		isFromTerm := other.isTerminal(from)

		//Anything from the other's initial goes from f's initial
		//Anything else gets offset
		if from != 0 {
			from += offset
		}

		if mappedFrom, ok := mapping[from]; ok {
			from = mappedFrom
		}

		if isFromTerm {
			f.addTerminal(from)
		}

		for ch, to := range transition {

			isToTerm := other.isTerminal(to)

			if to != 0 {
				to += offset
			}

			//Add transition if there isn't one already
			if existingTo, ok := f.Transitions[from][ch]; ok {
				mapping[to] = existingTo
			}

			//If we already have a mapping then use that
			if mappedTo, ok := mapping[to]; ok {
				to = mappedTo
			}

			if isToTerm {
				f.addTerminal(to)
			}

			f.AddTransition(from, to, []rune{ch})

			//Handle any internal loops by adding an additional state for them
			if existingTransitions, ok := f.Transitions[from]; ok {
				for over, existingTo := range existingTransitions {
					if from == existingTo {
						f.nextState++
						newState := f.nextState
						f.AddTransition(from, newState, []rune{over})
						f.AddTransition(newState, newState, []rune{over})
						f.addTerminal(newState)
					}
				}
			}
		}
	}
}

// Loop this automata on itself
func (f *FiniteState) Loop() {
	for from, transition := range f.Transitions {

		//If the transition comes from the initial state then we need
		//matching transitions from each terminal state
		if from == 0 {
			for _, state := range f.TerminalStates {
				for ch, to := range transition {
					f.AddTransition(state, to, []rune{ch})
				}
			}
		}
	}

	f.addTerminal(0)
}

// Negate this automata, i.e. make it non-accepting on it's original pattern
func (f *FiniteState) Negate() {
	terminals := []int{}

	for from, transition := range f.Transitions {
		if !f.isTerminal(from) {
			terminals = append(terminals, from)
		}

		for _, to := range transition {
			if !f.isTerminal(to) {
				terminals = append(terminals, to)
			}
		}
	}

	f.TerminalStates = terminals
}

func (f *FiniteState) String() string {
	sort.Ints(f.TerminalStates)
	str := fmt.Sprintf("Terminals: %v\n", f.TerminalStates)
	for from, transition := range f.Transitions {
		tran := ""
		for ch, to := range transition {
			tran += fmt.Sprintf("\n    %c => %d", ch, to)
		}
		str += fmt.Sprintf("%d: %s\n", from, tran)
	}
	return str
}

func (f *FiniteState) Execute(input string) bool {
	f.CurrentState = 0
	for _, ch := range input {
		if !f.consume(ch) {
			return false
		}
	}
	return f.isTerminal(f.CurrentState)
}

func (f *FiniteState) FindMatchEndIndex(input string) int {
	f.CurrentState = 0
	i := 0
	for _, ch := range input {
		if !f.consume(ch) {
			break
		}
		i++
	}

	if f.isTerminal(f.CurrentState) {
		return i
	}

	return 0
}

func (f *FiniteState) consume(ch rune) bool {
	return f.transition(f.CurrentState, ch)
}

func (f *FiniteState) transition(from int, ch rune) bool {

	if to, ok := f.Transitions[from][ch]; ok {
		f.CurrentState = to
		return true
	}

	return false
}

func (f *FiniteState) isTerminal(state int) bool {
	for _, val := range f.TerminalStates {
		if state == val {
			return true
		}
	}
	return false
}
