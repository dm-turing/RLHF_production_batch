package main

import (
	"fmt"
	"net/http"
)

// User struct to represent user input
type User struct {
	Name string
	Age  int
}

// ServeHTML function to generate HTML using fmt.Sprintf
func ServeHTML(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "Alice", // Replace with actual user input
		Age:  30,      // Replace with actual user input
	}

	htmlContent := fmt.Sprintf(`
	<html>
		<head><title>User Profile</title></head>
		<body>
			<h1>Hello, %s!</h1>
			<p>You are %d years old.</p>
		</body>
	</html>
	`, user.Name, user.Age)

	// Write the HTML content to the response
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, htmlContent)
}

func main() {
	http.HandleFunc("/", ServeHTML)
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
