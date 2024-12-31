package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

type Book struct {
	Title    string
	Author   string
	Genre    string
	ISBN     string
}

type User struct {
	Username string
	Salt     string // Added salt for salting passwords
	Password string
	Roles    []string // Added roles for authorization
}

const (
	RoleAdmin   = "admin"
	RoleUser    = "user"
)

var books []Book
var users []User

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Book Library Application!")
	initUsers()

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Login")
		fmt.Println("2. Exit")
		fmt.Print("Enter option: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			if user := login(scanner); user != nil {
				fmt.Println("\nLogged in successfully!")
				authorize(scanner, user)
			} else {
				fmt.Println("\nLogin failed.")
			}
		case "2":
			fmt.Println("\nExiting the application. Goodbye!")
			return
		default:
			fmt.Println("\nInvalid option. Please try again.")
		}
	}
}

func initUsers() {
	// Adding users with different roles for demonstration
	users = append(users, createUserWithRole("admin", "password", RoleAdmin))
	users = append(users, createUserWithRole("user", "password", RoleUser))
}

func createUserWithRole(username, password, role string) User {
	salt := base64.StdEncoding.EncodeToString(make([]byte, 32)) // Generate a 256-bit salt
	hash := sha256.New()
	hash.Write([]byte(password + salt)) // Hash the password with salt
	return User{
		Username: username,
		Salt:     salt,
		Password: hex.EncodeToString(hash.Sum(nil)),
		Roles:    []string{role},
	}
}

func login(scanner *bufio.Scanner) *User {
	// ... (Same implementation as before)
	// ...

	for _, user := range users {
		if user.Username == username && checkPassword(user.Salt, password, user.Password) {
			return &user
		}
	}

	return nil
}

func authorize(scanner *bufio.Scanner, user *User) {
	for {
		fmt.Println("\nAuthorized actions:")
		fmt.Println("1. Add a Book")
		fmt.Println("2. View Books")
		fmt.Println("3. Logout")
		fmt.Print("Enter action: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			if hasRole(user, RoleAdmin) {
				addBook()
			} else {
				fmt.Println("You don't have permission to add books.")
			}
		case "2":
			viewBooks()
		case "3":
			fmt.Println("\nLogging out.")
			return
		default:
			fmt.Println("\nInvalid action. Please try again.")
		}
	}
}

// Check if the user has a specific role
func hasRole(user *User, role string) bool {
	for _, userRole := range user.Roles {