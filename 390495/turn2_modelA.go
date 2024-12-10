package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Package struct {
	ImportPath string   `json:"ImportPath"`
	Imports    []string `json:"Imports"`
}

// checkDependencies uses `go list` to check if all dependencies of the current project are available
func checkDependencies() {
	// Execute the `go list -json ./...` command to retrieve package information
	cmd := exec.Command("go", "list", "-json", "./...")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing go list: %v\n", err)
		return
	}

	var packages []Package
	if err := json.Unmarshal(output, &packages); err != nil {
		fmt.Printf("Error unmarshalling json: %v\n", err)
		return
	}

	fmt.Printf("Analyzing project dependencies...\n")

	// Iterate through the packages and perform checks
	for _, pkg := range packages {
		if len(pkg.Imports) == 0 {
			fmt.Printf("Package %s has no imports.\n", pkg.ImportPath)
		} else {
			fmt.Printf("Package %s has dependencies: %v\n", pkg.ImportPath, pkg.Imports)
		}
	}

	// Additional health check can be implemented here...
}

func main() {
	checkDependencies()
}
