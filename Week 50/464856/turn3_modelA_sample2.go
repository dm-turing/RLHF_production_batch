package main

import (
	"fmt"
)

// Role represents a user role with associated permissions.
type Role struct {
	Name        string
	Permissions []string
}

// Permissions represents the set of permissions a user has.
type Permissions struct {
	Read  bool
	Write bool
}

func main() {
	roles := map[string]*Role{
		"admin":  {"admin", []string{"read", "write"}},
		"reader": {"reader", []string{"read"}},
	}

	// ... rest of the code
}

func checkPermission(username string, action string) (bool, error) {
	userRole := roles[username]
	if userRole == nil {
		return false, fmt.Errorf("User not found: %s", username)
	}

	for _, perm := range userRole.Permissions {
		if perm == action {
			return true, nil
		}
	}

	return false, fmt.Errorf("Permission denied for user: %s, action: %s", username, action)
}
