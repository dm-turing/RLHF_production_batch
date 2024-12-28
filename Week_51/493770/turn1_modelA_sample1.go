package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// Command interface
type Command interface {
	Execute(query url.Values) error
}

// ParseAgeCommand implementation of the Command interface
type ParseAgeCommand struct {
	name string
}

func NewParseAgeCommand(name string) *ParseAgeCommand {
	return &ParseAgeCommand{name: name}
}

func (c *ParseAgeCommand) Execute(query url.Values) error {
	ageStr, ok := query["age"]
	if !ok {
		return fmt.Errorf("age parameter missing")
	}
	age, err := strconv.Atoi(ageStr[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s is %d years old\n", c.name, age)
	return nil
}

// ParseEmailCommand implementation of the Command interface
type ParseEmailCommand struct {
	name string
}

func NewParseEmailCommand(name string) *ParseEmailCommand {
	return &ParseEmailCommand{name: name}
}

func (c *ParseEmailCommand) Execute(query url.Values) error {
	email, ok := query["email"]
	if !ok {
		return fmt.Errorf("email parameter missing")
	}
	fmt.Printf("%s's email is %s\n", c.name, email[0])
	return nil
}

func main() {
	query := url.Values{"age": []string{"25"}, "email": []string{"example@example.com"}}

	commands := []Command{
		NewParseAgeCommand("Alice"),
		NewParseEmailCommand("Bob"),
	}

	for _, cmd := range commands {
		if err := cmd.Execute(query); err != nil {
			fmt.Println(err)
		}
	}
}
