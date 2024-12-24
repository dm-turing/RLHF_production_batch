package main

import (
	"fmt"
)

// Role represents a role in the system
type Role string

// Permission represents a permission
type Permission string

// User represents a user in the system
type User struct {
	ID          int
	Name        string
	Role        Role
	Permissions []Permission
}

// CheckPermission checks if a user has a specific permission
func (u *User) CheckPermission(perm Permission) bool {
	for _, p := range u.Permissions {
		if p == perm {
			return true
		}
	}
	return false
}

func main() {
	// Define roles and permissions
	adminRole := Role("admin")
	userRole := Role("user")

	readPermission := Permission("read")
	writePermission := Permission("write")

	// Define users
	user1 := User{ID: 1, Name: "Alice", Role: userRole, Permissions: []Permission{readPermission}}
	user2 := User{ID: 2, Name: "Bob", Role: adminRole, Permissions: []Permission{readPermission, writePermission}}

	// Check permissions
	fmt.Println(user1.CheckPermission(readPermission))  // Output: true
	fmt.Println(user1.CheckPermission(writePermission)) // Output: false
	fmt.Println(user2.CheckPermission(readPermission))  // Output: true
	fmt.Println(user2.CheckPermission(writePermission)) // Output: true
}
