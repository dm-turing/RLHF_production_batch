package main

import (
	"io/ioutil"
	"log"
)

func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("error reading file %s: %v\n", filePath, err)
		return nil, err
	}
	return data, nil
}

func main() {
	readFile("config.json")
}
