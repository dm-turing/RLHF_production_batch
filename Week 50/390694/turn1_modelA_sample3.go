package main

import (
	"encoding/xml"
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string `xml:"name"`
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string `xml:"name"`
}

func (c Cat) Speak() string {
	return "Meow!"
}

func AnimalFactory(animalType string) Animal {
	switch animalType {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	default:
		return nil
	}
}

type AnimalXML struct {
	Type   string `xml:"type,attr"`
	Animal Animal `xml:",any"`
}

func main() {
	data := `
    <animals>
        <animal type="dog">
            <name>Buddy</name>
        </animal>
        <animal type="cat">
            <name>Whiskers</name>
        </animal>
    </animals>`

	var animals []Animal
	var animalsXML struct {
		Animals []AnimalXML `xml:"animal"`
	}

	if err := xml.Unmarshal([]byte(data), &animalsXML); err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	for _, aXML := range animalsXML.Animals {
		animal := AnimalFactory(aXML.Type)
		if animal != nil {
			xml.Unmarshal([]byte(aXML.Animal), &animal) // Unmarshal into the specific type
			animals = append(animals, animal)
		}
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
