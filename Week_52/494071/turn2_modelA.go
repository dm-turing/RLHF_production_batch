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
	City string `json:"city"`
}

var people = []Person{
	{ID: 1, Name: "Alice", Age: 30, City: "New York"},
	{ID: 2, Name: "Bob", Age: 25, City: "Los Angeles"},
	{ID: 3, Name: "Charlie", Age: 35, City: "Chicago"},
}

func personToMap(p Person) map[string]interface{} {
	return map[string]interface{}{
		"id":   p.ID,
		"name": p.Name,
		"age":  p.Age,
		"city": p.City,
	}
}

func peopleToMap(ps []Person) []map[string]interface{} {
	var pm []map[string]interface{}
	for _, p := range ps {
		pm = append(pm, personToMap(p))
	}
	return pm
}

// GET endpoint to fetch all people
func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsn, _ := json.MarshalIndent(peopleToMap(people), "", "  ")
	w.Write(jsn)
}

// GET endpoint to fetch a person by ID
func getPerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	personID, _ := strconv.Atoi(id)
	for _, person := range people {
		if person.ID == personID {
			w.Header().Set("Content-Type", "application/json")
			jsn, _ := json.MarshalIndent(personToMap(person), "", "  ")
			w.Write(jsn)
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
	w.Header().Set("Content-Type", "application/json")
	jsn, _ := json.MarshalIndent(personToMap(newPerson), "", "  ")
	w.Write(jsn)
}

func main() {
	http.HandleFunc("/api/people", getPeople)
	http.HandleFunc("/api/person", getPerson)
	http.HandleFunc("/api/add", addPerson)

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}
