package main

import (
	"crypto/bcrypt"
	"fmt"
)

// User represents a user with a username and hashed password.
type User struct {
	Username  string
	HashedPwd string
}

// VerifyPassword checks if the provided password matches the stored hashed password.
func (u *User) VerifyPassword(password string) bool {
	bytePass := []byte(password)
	byteHash := []byte(u.HashedPwd)
	return bcrypt.CompareHashAndPassword(byteHash, bytePass) == nil
}

func main() {
	users := map[string]*User{
		"admin": {"admin", "hashed_password_here"},
	}

	// ... rest of the code
}

func authenticateUser(username, password string) (bool, error) {
	if user, ok := users[username]; ok {
		if user.VerifyPassword(password) {
			return true, nil
		}
	}
	return false, fmt.Errorf("Authentication failed for user: %s", username)
}
