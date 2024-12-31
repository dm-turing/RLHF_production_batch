package main

import (
	"fmt"
)

// An interface for a function that takes no arguments and returns no value
type Action interface {
	Run()
}

func (a *Action) Run() {

}

// A function that takes an Action as an argument
func (a Action) executeAction() {
	a.Run()
}

func main() {
	// An anonymous function implementing the Action interface
	action := func() {
		fmt.Println("Running an action")
	}

	// Pass the action as a callback
	executeAction(action)
}
