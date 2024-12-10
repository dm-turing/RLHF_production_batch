package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// List outdated dependencies
	cmd := exec.Command("go", "list", "-u", "-m", "all")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error checking for outdated dependencies:", err)
		return
	}

	output := out.String()
	if output == "" {
		fmt.Println("No outdated dependencies found.")
	} else {
		fmt.Println("Outdated dependencies found:")
		fmt.Println(output)
	}
}
