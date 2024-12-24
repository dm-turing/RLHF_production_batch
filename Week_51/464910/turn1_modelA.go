package main

import (
	"fmt"
)

// Callback type
type stateCallback func(string)

// StateChanger represents an object that can change its state
type StateChanger struct {
	currentState string
	callbacks    []stateCallback
}

// NewStateChanger creates a new StateChanger instance
func NewStateChanger(initialState string) *StateChanger {
	return &StateChanger{
		currentState: initialState,
		callbacks:    []stateCallback{},
	}
}

// RegisterCallback registers a new callback for state changes
func (sc *StateChanger) RegisterCallback(cb stateCallback) {
	sc.callbacks = append(sc.callbacks, cb)
}

// ChangeState changes the current state and triggers callbacks
func (sc *StateChanger) ChangeState(newState string) {
	sc.currentState = newState
	for _, cb := range sc.callbacks {
		cb(sc.currentState)
	}
}

// NotifyStateChange is a sample callback function
func NotifyStateChange(state string) {
	fmt.Printf("State changed to: %s\n", state)
}

func main() {
	sc := NewStateChanger("Initial")
	sc.RegisterCallback(NotifyStateChange)

	sc.ChangeState("Running")
	sc.ChangeState("Stopped")
	sc.ChangeState("Running")
}
