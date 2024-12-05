package main

import (
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("go", "list", "-m", "-u")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	packages := strings.Split(string(output), "\n")
	for _, packge := range packages {
		if packge != "" {
			log.Printf("Update package: %s\n", packge)
			if err := exec.Command("go", "get", "-u", packge).Run(); err != nil {
				log.Printf("Error updating %s: %s\n", packge, err)
			}
		}
	}
}
