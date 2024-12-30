package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipCode"`
	Country   string `json:"country"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Customers []Customer

func importCSV(file string) (Customers, error) {
	var customers Customers
	f, err := os.Open(file)
	if err != nil {
		return customers, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	header, err := reader.Read()
	if err != nil {
		return customers, err
	}

	// Map CSV headers to struct field names
	fieldIndex := make(map[string]int)
	for idx, headerName := range header {
		fieldIndex[headerName] = idx
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return customers, err
		}

		customer := Customer{
			ID:        row[fieldIndex["id"]],
			Name:      row[fieldIndex["name"]],
			Email:     row[fieldIndex["email"]],
			Phone:     row[fieldIndex["phone"]],
			Address:   row[fieldIndex["address"]],
			City:      row[fieldIndex["city"]],
			State:     row[fieldIndex["state"]],
			ZipCode:   row[fieldIndex["zipCode"]],
			Country:   row[fieldIndex["country"]],
			CreatedAt: row[fieldIndex["createdAt"]],
			UpdatedAt: row[fieldIndex["updatedAt"]],
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func importJSON(file string) (Customers, error) {
	var customers Customers
	f, err := os.Open(file)
	if err != nil {
		return customers, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&customers)
	return customers, err
}

type CustomersXML struct {
	XMLName   xml.Name   `xml:"customers"`
	Customers []Customer `xml:"customer"`
}

func importXML(file string) (Customers, error) {
	var customersXML CustomersXML
	f, err := os.Open(file)
	if err != nil {
		return Customers{}, err
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	err = decoder.Decode(&customersXML)
	return customersXML.Customers, err
}

func exportCSV(file string, customers Customers) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	headers := []string{"id", "name", "email", "phone", "address", "city", "state", "zipCode", "country", "createdAt", "updatedAt"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	for _, customer := range customers {
		err = writer.Write([]string{customer.ID, customer.Name, customer.Email, customer.Phone, customer.Address, customer.City, customer.State, customer.ZipCode, customer.Country, customer.CreatedAt, customer.UpdatedAt})
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func exportJSON(file string, customers Customers) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	return encoder.Encode(customers)
}

func exportXML(file string, customers Customers) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	customersXML := CustomersXML{Customers: customers}
	encoder := xml.NewEncoder(f)
	encoder.Indent("", "  ")
	return encoder.Encode(customersXML)
}

func main() {
	c := Customer{
		ID:        "400",
		Name:      "Mage",
		Email:     "mage@dota2@.com",
		Phone:     "123-456-7890",
		Address:   "Radiant",
		City:      "city4",
		State:     "state4",
		ZipCode:   "100123",
		CreatedAt: "30-12-2024",
		UpdatedAt: "30-12-2024",
	}
	customers := [...]Customer{c}

	importCSV("data.csv")
	importXML("data.xml")
	importJSON("data.json")
	exportCSV("data1.csv", customers[:])
	exportXML("data1.xml", customers[:])
	exportJSON("data1.json", customers[:])
}
