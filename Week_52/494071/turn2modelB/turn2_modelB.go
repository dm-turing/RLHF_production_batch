package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

var people = []Person{
	{ID: 1, Name: "Alice", Age: 30, Address: &Address{Street: "street1", City: "city1", ZipCode: "100232"}},
	{ID: 2, Name: "Bob", Age: 25, Address: &Address{Street: "street2", City: "city2", ZipCode: "102232"}},
	{ID: 3, Name: "Charlie", Age: 35, Address: &Address{Street: "street3", City: "city3", ZipCode: "120232"}},
}

// Complex Nested Response using map
func getComplexPerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	for _, person := range people {
		if person.ID == parseInt(id) {
			address, err := json.Marshal(person.Address)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := fmt.Sprintf(`{"id":%d,"name":"%s","age":%d,"address":%s}`, person.ID, person.Name, person.Age, string(address))
			_, err = w.Write([]byte(response))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

// rest of the code remains the same
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
	http.HandleFunc("/api/people", getComplexPerson)
	http.HandleFunc("/api/person", getPerson)
	http.HandleFunc("/api/add", addPerson) // POST method overrides GET method for the same path

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
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
