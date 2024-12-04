package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// Start the pprof server
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Your application logic here
}
