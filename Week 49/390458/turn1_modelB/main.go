package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/example/models"
	"github.com/example/xmlutils"
)

func main() {
	file, err := os.Open("example.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data models.RSS
	err = xml.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	fmt.Println("Parsed XML:")
	xmlutils.PrettyPrint(data)
}
