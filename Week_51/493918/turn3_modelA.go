package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type timeBasedCache struct {
	cache map[string]interface{}
	ttl   time.Duration
	mu    sync.RWMutex
}

var lastinserted time.Time

func (c *timeBasedCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	lastinserted = time.Now()
	c.cache[key] = value
	c.cleanup()
}

func (c *timeBasedCache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// now := time.Now()
	if value, ok := c.cache[key]; ok {
		// For demonstration, we simply clean the cache instead of checking individual TTL
		c.cleanup()
		return value
	}
	return nil
}

func (c *timeBasedCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys := make([]string, 0, len(c.cache))
	for k := range c.cache {
		keys = append(keys, k)
	}

	for _, key := range keys {
		// Remove all keys before a certain time
		if strings.Contains(key, time.Now().Sub(lastinserted).String()) {
			delete(c.cache, key)
		}
	}
}

func main() {
	r := mux.NewRouter()
	c := timeBasedCache{cache: make(map[string]interface{}), ttl: 1 * time.Second}
	r.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		value := r.FormValue("value")

		if key != "" {
			if value != "" {
				c.Set(key, value)
			}

			response, exists := c.Get(key).(string)
			if exists {
				fmt.Fprintf(w, "Cache Hit: %s\n", response)
			} else {
				fmt.Fprintf(w, "Cache Miss: %s\n", key)
			}
		}
	}).Methods("GET", "POST")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
