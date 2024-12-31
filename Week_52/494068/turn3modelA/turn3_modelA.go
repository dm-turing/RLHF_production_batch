package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Genre  string
	ISBN   string
}

type User struct {
	Username string
	Salt     string // Added salt for salting passwords
	Password string
	Role     string // Added role for role-based authorization
}

var books []Book
var users []User

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Book Library Application!")

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add a Book")
		fmt.Println("2. View Books")
		fmt.Println("3. Login")
		fmt.Println("4. Exit")
		fmt.Print("Enter option: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			addBook()
		case "2":
			viewBooks()
		case "3":
			if login(scanner) {
				fmt.Println("\nLogged in successfully!")
				authorize(scanner)
			} else {
				fmt.Println("\nLogin failed.")
			}
		case "4":
			fmt.Println("\nExiting the application. Goodbye!")
			return
		default:
			fmt.Println("\nInvalid option. Please try again.")
		}
	}
}

func addBook() {
	var title, author, genre, isbn string
	fmt.Print("Enter book title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter book author: ")
	fmt.Scanln(&author)
	fmt.Print("Enter book genre: ")
	fmt.Scanln(&genre)
	fmt.Print("Enter book ISBN: ")
	fmt.Scanln(&isbn)

	books = append(books, Book{Title: title, Author: author, Genre: genre, ISBN: isbn})
	fmt.Println("Book added successfully!")
}

func viewBooks() {
	if len(books) == 0 {
		fmt.Println("No books in the library.")
		return
	}

	fmt.Println("Books in the library:")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Genre: %s, ISBN: %s\n", book.Title, book.Author, book.Genre, book.ISBN)
	}
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
			return checkPassword(user.Salt, password, user.Password) && hasRole(user.Role, "viewer", "admin")
		}
	}

	return false
}

func checkPassword(salt, provided, hashed string) bool {
	// Combine salt with provided password
	combined := strings.Join([]string{provided, salt}, "")
	// Generate new hash of the combined string
	newHash := sha256.New()
	newHash.Write([]byte(combined))
	newHashed := hex.EncodeToString(newHash.Sum(nil))
	return newHashed == hashed
}

func hasRole(roles ...string) bool {
	userRole := roles[0]
	requiredRole := roles[1:]
	for _, role := range requiredRole {
		if userRole == role {
			return true
		}
	}
	return false
}

func authorize(scanner *bufio.Scanner) {
	for {
		fmt.Println("\nAuthorized actions:")
		fmt.Println("1. Add a Book")
		fmt.Println("2. View Books")
		fmt.Println("3. Logout")
		fmt.Print("Enter action: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			if hasRole(currentUser().Role, "admin") {
				addBook()
			} else {
				fmt.Println("Insufficient permissions. Please contact an admin.")
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

func currentUser() User {
	// Placeholder for a more comprehensive user management system.
	for _, user := range users {
		if user.Username == "admin" {
			return user
		}
	}
	return User{} // Empty user in case of no logins.
}

// Simple function to initialize the system
func init() {
	// Adding a user with hashed password for demonstration
	salt := base64.StdEncoding.EncodeToString(make([]byte, 32)) // Generate a 256-bit salt
	hash := sha256.New()
	hash.Write([]byte("password" + salt)) // Hash the password with salt
	users = append(users, User{Username: "admin", Salt: salt, Password: hex.EncodeToString(hash.Sum(nil)), Role: "viewer"})
}
