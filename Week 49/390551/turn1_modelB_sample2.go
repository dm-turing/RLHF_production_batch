package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
)

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	return file, nil
}

func readFile(filename string) ([]byte, error) {
	file, err := openFile(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}
	return data, nil
}

func main() {
	data, err := readFile("non_existent_file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
