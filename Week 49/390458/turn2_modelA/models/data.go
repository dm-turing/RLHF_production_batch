package models

type Person struct {
	Name    string `xml:"name"`
	Age     int    `xml:"age"`
	Address string `xml:"address"`
}

type People struct {
	Persons []Person `xml:"person"`
}
