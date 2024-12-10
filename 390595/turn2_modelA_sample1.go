package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func lockFile(file *os.File) error {
	return unix.Flock(int(file.Fd()), unix.LOCK_EX)
}

func unlockFile(file *os.File) error {
	return unix.Flock(int(file.Fd()), unix.LOCK_UN)
}

func writeToFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := lockFile(file); err != nil {
		return err
	}
	defer unlockFile(file) // Ensure the file is unlocked when done

	_, err = file.WriteString(content + "\n")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	filePath := "example.txt"

	err := writeToFile(filePath, "Writing from goroutine")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Write successful.")
	}
}
