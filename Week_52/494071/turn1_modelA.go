package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people = []Person{
	{ID: 1, Name: "Alice", Age: 30},
	{ID: 2, Name: "Bob", Age: 25},
	{ID: 3, Name: "Charlie", Age: 35},
}

// GET endpoint to fetch all people
func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GET endpoint to fetch a person by ID
func getPerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	for _, person := range people {
		if person.ID == parseInt(id) {
			json.NewEncoder(w).Encode(person)
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

// POST endpoint to add a new person
func addPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson Person
	if err := json.NewDecoder(r.Body).Decode(&newPerson); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	newPerson.ID = len(people) + 1
	people = append(people, newPerson)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func main() {
	http.HandleFunc("/api/people", getPeople)
	http.HandleFunc("/api/person", getPerson)
	http.HandleFunc("/api/person", addPerson) // POST method overrides GET method for the same path

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
