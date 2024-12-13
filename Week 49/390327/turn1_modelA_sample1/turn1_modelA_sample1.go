package main

import "fmt"

type Config struct {
	APIToken string
}

type ActionPerformer interface {
	PerformAction() string
}

type MyAction struct {
	config Config
}

func (a *MyAction) PerformAction() string {
	return fmt.Sprintf("Performing action with token: %s", a.config.APIToken)
}

// Refactored function
func DoSomething(performer ActionPerformer) {
	fmt.Println(performer.PerformAction())
}

// Legacy code with global state
var config = Config{APIToken: "123"}

func PerformAction() {
	// uses global config
	fmt.Println("Performing action with token: ", config.APIToken)
}

func main() {
	DoSomething(&MyAction{config: config})
}
