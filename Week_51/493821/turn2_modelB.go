package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define the transifex project, resource, and language codes
	project := "your-project"
	resource := "your-resource"
	languages := []string{"en", "fr", "es"} // Add more language codes as needed

	// Fetch translations from transifex
	for _, language := range languages {
		cmd := exec.Command("tx", "pull", fmt.Sprintf("-p %s -r %s -l %s", project, resource, language))
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error fetching translations for", language, ":", err)
			fmt.Println(string(output))
			os.Exit(1)
		}
	}

	// Generate translation files (e.g., PO)
	for _, language := range languages {
		cmd := exec.Command("tx", "po", fmt.Sprintf("-p %s -r %s -l %s", project, resource, language))
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error generating PO files for", language, ":", err)
			fmt.Println(string(output))
			os.Exit(1)
		}
	}

	// Run automated tests for translations
	cmd := exec.Command("go", "test", "-v", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Translation tests failed:", err)
		fmt.Println(string(output))
		os.Exit(1)
	}

	// Create deployment package (e.g., zip)
	// Add your deployment package creation logic here

	fmt.Println("Translations successfully fetched, generated, and tested.")
}
