package main

import (
	"fmt"
)

// Callback types
type messageReceivedCallback func(string, string)
type userJoinedCallback func(string)
type userLeftCallback func(string)
type channelJoinedCallback func(string)
type channelLeftCallback func(string)

// User represents a user in the chat application
type User struct {
	name              string
	messageReceivedCB messageReceivedCallback
	userJoinedCB      userJoinedCallback
	userLeftCB        userLeftCallback
}

// Channel represents a chat channel in the application
type Channel struct {
	name              string
	users             map[string]*User
	messageReceivedCB messageReceivedCallback
	userJoinedCB      userJoinedCallback
	userLeftCB        userLeftCallback
}

// ChatApplication represents the entire chat system
type ChatApplication struct {
	channels     map[string]*Channel
	userJoinedCB userJoinedCallback
	userLeftCB   userLeftCallback
}

// NewUser creates a new user
func NewUser(name string) *User {
	return &User{name: name}
}

// NewChannel creates a new chat channel
func NewChannel(name string) *Channel {
	return &Channel{name: name, users: make(map[string]*User)}
}

// NewChatApplication creates a new chat application
func NewChatApplication() *ChatApplication {
	return &ChatApplication{channels: make(map[string]*Channel)}
}

// RegisterMessageReceivedCallback registers a callback for message received events
func (u *User) RegisterMessageReceivedCallback(cb messageReceivedCallback) {
	u.messageReceivedCB = cb
}

// RegisterUserJoinedCallback registers a callback for user joined events
func (u *User) RegisterUserJoinedCallback(cb userJoinedCallback) {
	u.userJoinedCB = cb
}

// RegisterUserLeftCallback registers a callback for user left events
func (u *User) RegisterUserLeftCallback(cb userLeftCallback) {
	u.userLeftCB = cb
}

// RegisterMessageReceivedCallback registers a callback for message received events
func (c *Channel) RegisterMessageReceivedCallback(cb messageReceivedCallback) {
	c.messageReceivedCB = cb
}

// RegisterUserJoinedCallback registers a callback for user joined events
func (c *Channel) RegisterUserJoinedCallback(cb userJoinedCallback) {
	c.userJoinedCB = cb
}

// RegisterUserLeftCallback registers a callback for user left events
func (c *Channel) RegisterUserLeftCallback(cb userLeftCallback) {
	c.userLeftCB = cb
}

// RegisterUserJoinedCallback registers a callback for user joined events
func (ca *ChatApplication) RegisterUserJoinedCallback(cb userJoinedCallback) {
	ca.userJoinedCB = cb
}

// RegisterUserLeftCallback registers a callback for user left events
func (ca *ChatApplication) RegisterUserLeftCallback(cb userLeftCallback) {
	ca.userLeftCB = cb
}

// HandleMessageReceived is a callback function that handles message received events
func HandleMessageReceived(user, message string) {
	fmt.Printf("%s: %s\n", user, message)
}

// HandleUserJoined is a callback function that handles user joined events
func HandleUserJoined(user string) {
	fmt.Printf("%s joined the channel.\n", user)
}

// HandleUserLeft is a callback function that handles user left events
