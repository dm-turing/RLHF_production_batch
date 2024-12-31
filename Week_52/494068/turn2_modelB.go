package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
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
	Password []byte
	Salt     []byte
}

var books []Book
var users []User

const passwordHashBytes = 64
const saltBytes = 32

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Book Library Application!")

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Enter option: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			if err := register(scanner); err != nil {
				fmt.Println("Registration failed:", err)
			} else {
				fmt.Println("Registration successful!")
			}
		case "2":
			if login(scanner) {
				fmt.Println("\nLogged in successfully!")
				authorize(scanner)
			} else {
				fmt.Println("\nLogin failed.")
			}
		case "3":
			fmt.Println("\nExiting the application. Goodbye!")
			return
		default:
			fmt.Println("\nInvalid option. Please try again.")
		}
	}
} 
 
func register(scanner *bufio.Scanner) error {
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	salt, err := generateRandomSalt(saltBytes)
	if err != nil {
		return err
	}

	hashedPassword, err := hashPassword(password, salt)
	if err != nil {
		return err
	}

	users = append(users, User{Username: username, Password: hashedPassword, Salt: salt})
	return nil
}

func login(scanner *bufio.Scanner) bool {
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	for _, user := range users {
		if user.Username == username {
			if err := verifyPassword(password, user.Password, user.Salt); err == nil {
				return true
			}
		}
	}

	return false
}

func authorize(scanner *bufio.Scanner) {
	// ... (Same as before)
}
 
func generateRandomSalt(bytes int) ([]byte, error) {
	salt := make([]byte, bytes)
	_, err := rand.Read(salt)
	return salt, err
}

func hashPassword(password string, salt []byte) ([]byte, error) {
	hash := sha512.New()
	hash.Write(salt)
	hash.Write([]byte(password))
	return hash.Sum(nil), nil
}

func verifyPassword(password string, hashedPassword []byte, salt []byte) error {
	newHashedPassword, err := hashPassword(password, salt)
	if err != nil {
		return err
	}
	
	if !equal(hashedPassword, newHashedPassword) {
		return errors.New("password mismatch")