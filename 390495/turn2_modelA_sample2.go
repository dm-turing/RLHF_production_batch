package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Run 'go mod tidy' to remove unused dependencies
	cmd := exec.Command("go", "mod", "tidy")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running 'go mod tidy':", err)
		return
	}

	fmt.Println("Unused dependencies removed from go.mod.")
}
