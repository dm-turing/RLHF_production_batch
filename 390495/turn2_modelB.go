package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Tidy the Go module to ensure all dependencies are up-to-date
	if err := tidyModules(); err != nil {
		log.Fatalf("failed to tidy modules: %v", err)
	}

	// Check for vulnerabilities
	if err := checkVulnerabilities(); err != nil {
		log.Fatalf("failed to check for vulnerabilities: %v", err)
	}

	fmt.Println("Project health check complete. No critical issues found.")
}

func tidyModules() error {
	cmd := exec.CommandContext(context.Background(), "go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkVulnerabilities() error {
	output, err := exec.CommandContext(context.Background(), "govulncheck", "-mod=readonly").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run govulncheck: %v", err)
	}

	if strings.Contains(string(output), "Vulnerabilities:") {
		fmt.Println("Vulnerabilities found:")
		fmt.Println(string(output))
		return fmt.Errorf("vulnerabilities detected")
	}

	return nil
}
