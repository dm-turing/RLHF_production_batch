package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type File struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type Worker struct {
	ID    int
	Files map[string]File
}

type Master struct {
	Workers      []Worker
	LoadBalancer func(int) int // Random load balancer
}

func (m *Master) StoreFile(w http.ResponseWriter, r *http.Request) {
	file := File{
		Name: "example",
		Data: "There is some sample data here",
	}
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	workerID := m.LoadBalancer(len(m.Workers))
	m.Workers[workerID].Files[file.Name] = file
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "File %s stored successfully\n", file.Name)
}

func (m *Master) RetrieveFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("name")
	if fileName == "" {
		http.Error(w, "File name not provided", http.StatusBadRequest)
		return
	}

	workerID := m.LoadBalancer(len(m.Workers))
	file, ok := m.Workers[workerID].Files[fileName]
	if !ok {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(file.Data))
}

func main() {
	m := Master{
		Workers: []Worker{
			{ID: 0, Files: make(map[string]File)},
			{ID: 1, Files: make(map[string]File)},
		},
		LoadBalancer: func(n int) int {
			return rand.Intn(n)
		},
	}

	http.HandleFunc("/store", m.StoreFile)
	http.HandleFunc("/retrieve", m.RetrieveFile)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
