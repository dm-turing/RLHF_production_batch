package main

import (
	"html/template"
	"log"
	"net/http"
)

type User1 struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User1{Name: r.URL.Query().Get("name")}
	tmpl, err := template.ParseFiles("welcome.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	err = tmpl.Execute(w, user.Name)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
