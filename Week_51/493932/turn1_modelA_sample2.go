package main

import (
	"fmt"
)

type Config struct {
	Host string `xml:"host"`
	Port int    `xml:"port"`
}

func main() {
	config := Config{Host: "localhost", Port: 8080}
	xmlString := fmt.Sprintf(`<config><host>%s</host><port>%d</port></config>`, config.Host, config.Port)
	fmt.Println(xmlString)
}
