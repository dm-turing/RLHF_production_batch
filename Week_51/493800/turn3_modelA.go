package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type FileData struct {
	Name      string    `json:"name"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

type Master struct {
	Metadata map[string]FileData
	Slaves   []string
	Mu       sync.Mutex
	// Additional details for consistency and failover would go here
}

func (m *Master) SaveFile(w http.ResponseWriter, r *http.Request) {
	fileData := FileData{}
	err := json.NewDecoder(r.Body).Decode(&fileData)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	m.Mu.Lock()
	m.Metadata[fileData.Name] = fileData
	m.Mu.Unlock()

	// Send the updated data to all slaves
	for _, slave := range m.Slaves {
		sendDataToSlave(slave, fileData)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "File %s saved successfully\n", fileData.Name)
}

func sendDataToSlave(slave string, fileData FileData) {
	b, _ := json.Marshal(fileData)
	resp, err := http.Post(fmt.Sprintf("http://%s/save-file", slave), "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Error sending data to slave %s: %v\n", slave, err)
	}
	defer resp.Body.Close()
}

func main() {
	m := Master{Metadata: make(map[string]FileData), Slaves: []string{"slave1", "slave2"}}

	http.HandleFunc("/save-file", m.SaveFile)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
