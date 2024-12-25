package main

import (
	"fmt"
	"sync"
)

type State struct {
	value int
	lock  sync.Mutex
}
type Callback func()
type StateRollback struct {
	callbacks []Callback
}

func (s *State) Set(value int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value = value
}
func (s *State) Get() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.value
}
func (sr *StateRollback) AddCallback(cb Callback) {
	sr.callbacks = append(sr.callbacks, cb)
}
func (sr *StateRollback) Rollback() {
	for i := len(sr.callbacks) - 1; i >= 0; i-- {
		sr.callbacks[i]()
	}
}
func Step1(state *State, sr *StateRollback) {
	fmt.Println("Step 1: Adding 10 to state")
	state.Set(state.Get() + 10)
	sr.AddCallback(func() {
		fmt.Println("Step 1: Rolling back")
		state.Set(state.Get() - 10)
	})
}
func Step2(state *State, sr *StateRollback) {
	fmt.Println("Step 2: Multiplying state by 2")
	state.Set(state.Get() * 2)
	sr.AddCallback(func() {
		fmt.Println("Step 2: Rolling back")
		state.Set(state.Get() / 2)
	})
}
func Step3(state *State, sr *StateRollback) {
	fmt.Println("Step 3: Adding 5 to state")
	state.Set(state.Get() + 5)
	sr.AddCallback(func() {
		fmt.Println("Step 3: Rolling back")
		state.Set(state.Get() - 5)
	})
}
func main() {
	state := &State{value: 0}
	sr := &StateRollback{}

	Step1(state, sr)
	Step2(state, sr)
	Step3(state, sr)

	fmt.Println("Final state:", state.Get())

	fmt.Println("\nPerforming rollback:")
	sr.Rollback()

	fmt.Println("Final state after rollback:", state.Get())
}
