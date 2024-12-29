package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	file    string
	help    bool
	version int
)

func main() {
	flag.StringVar(&file, "file", "", "file to version")
	flag.BoolVar(&help, "help", false, "show help")
	flag.IntVar(&version, "version", 0, "retrieve a specific version (0 for latest)")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if file == "" {
		fmt.Println("No file specified")
		flag.Usage()
		return
	}
	_, err := os.Stat(file)
	switch err == nil {
	case false:
		fmt.Printf("File '%s' does not exist\n", file)
	default:
		if version != 0 {
			restoreFile(file, version)
			return
		}

		versionName := getCurrentVersion(file)
		backup(file, versionName)
		fmt.Printf("File '%s' saved at version %d\n", file, versionName)
	}
}

func getCurrentVersion(file string) int {
	versionName, err := ioutil.ReadFile(fmt.Sprintf("%s.version", file))
	if err != nil {
		return 0 // Initialize to 0 if no version file exists
	}
	return int(versionName)
}

func backup(file string, version int) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Printf("Failed to get file info for '%s': %v\n", file, err)
		return
	}

	newVersion := version + 1
	newVersionName := fmt.Sprintf("%d", newVersion)

	if err := ioutil.WriteFile(fmt.Sprintf("%s.%d", file, newVersion), fileInfo.Mode(), fileInfo.Sys()); err != nil {
		fmt.Printf("Failed to back up file '%s': %v\n", file, err)
		return
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%s.version", file), []byte(newVersionName), 0644); err != nil {
		fmt.Printf("Failed to update version file: %v\n", err)
		return
	}
}

func restoreFile(file string, version int) {
	err := ioutil.WriteFile(file, []byte(""), 0644)
	if err != nil {
		fmt.Printf("Failed to clear original file '%s': %v\n", file, err)
		return
	}

	versionName := fmt.Sprintf("%d", version)
	byteSource, err := ioutil.ReadFile(fmt.Sprintf("%s.%s", file, versionName))
	if err != nil {
		fmt.Printf("Failed to read versioned file '%s.%s': %v\n", file, versionName, err)
		return
	}

	if err := ioutil.WriteFile(file, byteSource, 0644); err != nil {
		fmt.Printf("Failed to write versioned file back to '%s': %v\n", file, err)
		return
	}
}
