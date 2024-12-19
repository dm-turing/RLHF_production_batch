package main

import (
	"fmt"
	"os"
)

func main() {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
	}
	fmt.Println("Server Host:", host)
}
