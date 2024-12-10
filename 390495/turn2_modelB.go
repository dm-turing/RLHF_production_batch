package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// PackageInfo represents information about a Go package.
type PackageInfo struct {
	ImportPath  string   `json:"ImportPath"`
	Dir         string   `json:"Dir"`
	GoFiles     []string `json:"GoFiles"`
	TestGoFiles []string `json:"TestGoFiles"`
	Imports     []string `json:"Imports"`
	Error       *string  `json:"Error"`
}

func main() {
	// List all packages in the current directory and subdirectories
	cmd := exec.Command("go", "list", "-json", "./...")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run go list: %v", err)
	}

	var packages []PackageInfo
	if err := json.Unmarshal(out.Bytes(), &packages); err != nil {
		log.Fatalf("Failed to parse JSON output: %v", err)
	}

	fmt.Println("Project Health Analysis:")
	fmt.Println("------------------------")

	for _, pkg := range packages {
		if pkg.Error != nil {
			fmt.Printf("🔴 Error in package %s: %s\n", pkg.ImportPath, *pkg.Error)
		} else {
			fmt.Printf("✅ Package %s (Dir: %s) is buildable.\n", pkg.ImportPath, pkg.Dir)
			fmt.Printf("  - Go Files: %v\n", pkg.GoFiles)
			fmt.Printf("  - Test Go Files: %v\n", pkg.TestGoFiles)
			fmt.Printf("  - Dependencies: %v\n", pkg.Imports)
		}
	}
}
