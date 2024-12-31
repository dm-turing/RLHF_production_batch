package main

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
)

// Book struct represents a book in the library
type Book struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	ISBN    string `json:"isbn"`
	OwnerID uint   `json:"owner_id"` // Foreign key to User
}

// User struct represents a user in the application
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *gorm.DB

func initializeDatabase() {
	var err error
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{}, &User{})
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHashedPwd := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHashedPwd, plainPwd)
	return err == nil
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = hashAndSalt([]byte(user.Password))
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var dbUser User
	db.Where("username = ?", user.Username).Find(&dbUser)
	if dbUser.ID == 0 {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	if !comparePasswords(dbUser.Password, []byte(user.Password)) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	// Generate JWT token and send it in response
	tokenString := generateJWT(dbUser.ID)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from JWT token
	userID, ok := getUserIDFromToken(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book.OwnerID = userID
	db.Create(&book)
	json.NewEncoder(w).Encode(book)
}
