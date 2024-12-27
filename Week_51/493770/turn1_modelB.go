package main

import (
	"fmt"
	"net/url"
)

// Command is an interface defining the operation to be performed on the query values.
type Command interface {
	Execute(values url.Values)
}

// FetchAuthorCommand fetches the author from the query.
type FetchAuthorCommand struct{}

// Execute implementation of Command interface for FetchAuthorCommand.
func (f *FetchAuthorCommand) Execute(values url.Values) {
	if author := values.Get("author"); author != "" {
		fmt.Println("Author:", author)
	} else {
		fmt.Println("Author not found.")
	}
}

// FetchTitleCommand fetches the title from the query.
type FetchTitleCommand struct{}

// Execute implementation of Command interface for FetchTitleCommand.
func (f *FetchTitleCommand) Execute(values url.Values) {
	if title := values.Get("title"); title != "" {
		fmt.Println("Title:", title)
	} else {
		fmt.Println("Title not found.")
	}
}

// Invoker stores the command that will be executed.
type Invoker struct {
	command Command
}

// SetCommand sets the command that will be executed by the invoker.
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteInvoker executes the previously set command.
func (i *Invoker) ExecuteInvoker() {
	i.command.Execute(nil) // Assuming query values are set elsewhere.
}

func main() {
	// Create a URL query value
	queryValues := url.Values{}
	queryValues.Set("author", "Neil Gaiman")
	queryValues.Set("title", "American Gods")

	//Creating  invoker and setting different commands
	invoker := new(Invoker)

	// First time, fetch author using command pattern
	invoker.SetCommand(&FetchAuthorCommand{})
	invoker.ExecuteInvoker() // Output: Author: Neil Gaiman

	//Second time, fetch title using command pattern
	invoker.SetCommand(&FetchTitleCommand{})
	invoker.ExecuteInvoker() // Output: Title: American Gods
}
