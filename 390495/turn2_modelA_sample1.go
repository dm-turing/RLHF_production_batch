package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Run 'go mod tidy' to clean up go.mod
	cmd := exec.Command("go", "mod", "tidy")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running 'go mod tidy':", err)
		return
	}

	// Run 'go build' to check for missing dependencies
	cmd = exec.Command("go", "build", "./...")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Some dependencies are missing or there are other build errors.")
		return
	}

	fmt.Println("All dependencies are resolved and buildable.")
}
