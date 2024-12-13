package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func openFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to open file %s", filename))
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()
	return nil
}

func main() {
	if err := openFile("non_existent_file.txt"); err != nil {
		fmt.Println(err)
	}
}
