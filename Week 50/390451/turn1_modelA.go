package main

import (
	"fmt"
	"os"
)

func main() {
	cloudProvider := os.Getenv("CLOUD_PROVIDER")
	switch cloudProvider {
	case "aws":
		// AWS specific configuration
		fmt.Println("For AWS deployments")
	case "azure":
		// Azure specific configuration
		fmt.Println("For Azure deployments")
	case "gcp":
		// GCP specific configuration
		fmt.Println("For GCP deployments")
	}
}
