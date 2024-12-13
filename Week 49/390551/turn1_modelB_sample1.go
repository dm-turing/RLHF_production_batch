package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func readFile(filename string) ([]byte, error) {
	file, err := openFile(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
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
