package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("name")
	htmlContent := fmt.Sprintf(`
        <html>
            <head>
                <title>Welcome</title>
            </head>
            <body>
                <h1>Hello, %s!</h1>
                <p>Welcome to our website.</p>
            </body>
        </html>
    `, userName)
	w.Write([]byte(htmlContent))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
