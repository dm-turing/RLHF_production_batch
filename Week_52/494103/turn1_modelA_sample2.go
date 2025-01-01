package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrapf(err, "failed to open file %s", filename)
	}
	defer file.Close()
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	err := readFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", errors.Wrapf(err, "error reading file %s", filename))
		os.Exit(1)
	}

	fmt.Println("File read successfully.")
}
