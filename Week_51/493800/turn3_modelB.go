package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type File struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type Worker struct {
	ID       int
	Files    map[string]File
	mutex    sync.RWMutex // For thread safety
	isAlive  bool
	lastBeat time.Time
}

type Master struct {
	Workers []Worker
	heartbeatInterval time.Duration
}

func newMaster(heartbeatInterval time.Duration) *Master {
	return &Master{
		Workers: make([]Worker, 0),
		heartbeatInterval: heartbeatInterval,
	}
}

func (m *Master) registerWorker(w Worker) {
	m.Workers = append(m.Workers, w)
}

func (m *Master) heartbeatMonitor() {
	ticker := time.NewTicker(m.heartbeatInterval)
	for range ticker.C {
		for i, worker := range m.Workers {
			if time.Since(worker.lastBeat) > m.heartbeatInterval*2 {
				log.Printf("Worker %d is dead, marking it as failed\n", worker.ID)
				worker.mutex.Lock()
				worker.isAlive = false
				worker.mutex.Unlock()
			}
		}
	}
}

func (w *Worker) heartbeat(m *Master) {
	ticker := time.NewTicker(m.heartbeatInterval)
	for range ticker.C {
		w.mutex.Lock()
		w.isAlive = true
		w.lastBeat = time.Now()
		w.mutex.Unlock()
	}
}

func (m *Master) StoreFile(w http.ResponseWriter, r *http.Request) {
	file := File{}
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}
	
	// Perform 3-way replication
	replicas := m.getLiveWorkers(3)
	if len(replicas) < 3 {
		http.Error(w, "Not enough live workers for replication", http.StatusServiceUnavailable)
		return
	}
	
	var wg sync.WaitGroup
	wg.Add(3)
	
	for _, replica := range replicas {
		go func(replica Worker) {
			replica.mutex.Lock()
			replica.Files[file.Name] = file
			replica.mutex.Unlock()
			wg.Done()
		}(replica)
	}
	
	wg.Wait()
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "File %s stored successfully\n", file.Name)
}

func (m *Master) RetrieveFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("name")
	if fileName == "" {
		http.Error(w, "File name not provided", http.StatusBadRequest)
		return
	}
	
	// Perform read quorum (2/3 majority)
	replicas := m.getLiveWorkers(3)
	if len(replicas) < 3 {
		http.Error(w, "Not enough live workers for read quorum", http.StatusServiceUnavailable)
		return
	}
	
	var wg sync.WaitGroup
	wg.Add(3)
	
	var fileData string
	var fileFound bool
	var errors []error
	
	for _, replica := range replicas {
		go func(replica Worker) {
			replica.mutex.RLock()
			file, ok := replica.Files[fileName]
			replica.mutex.RUnlock()