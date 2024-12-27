package main

import (
	"errors"
	"fmt"
)

type CredentialError struct {
	Username string
	Retry    func()
}

func (e *CredentialError) Error() string {
	return fmt.Sprintf("credential error for user %s", e.Username)
}

func login() (*CredentialError, error) {
	// Login logic
	username := "user"
	return &CredentialError{Username: username, Retry: func() { login() }}, errors.New("failed to login")
}

func main() {
	c, err := login()
	fmt.Println(c.Error(), err.Error())
}
