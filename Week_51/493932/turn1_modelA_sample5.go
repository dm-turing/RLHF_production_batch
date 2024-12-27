package main

import (
	"encoding/xml"
	"fmt"
)

type Config struct {
	XMLName xml.Name `xml:"config"`
	Host    string   `xml:"host"`
	Port    int      `xml:"port"`
}

func main() {
	config := Config{Host: "localhost", Port: 8080}
	xmlBytes, err := xml.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(xmlBytes))
}
