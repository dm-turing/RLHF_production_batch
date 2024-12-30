package main

import (
    "encoding/csv"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io"
    "log"
    "os"
)

// Customer struct represents a customer record
type Customer struct {
    ID        string `json:"id" xml:"id"`
    Name      string `json:"name" xml:"name"`
    Email     string `json:"email" xml:"email"`
    Phone     string `json:"phone" xml:"phone"`
    Address   string `json:"address" xml:"address"`
}

// Importer interface defines the import method
type Importer interface {
    Import(io.Reader) ([]Customer, error)
}

// Exporter interface defines the export method
type Exporter interface {
    Export(io.Writer, []Customer) error
}

// CSVImporter implements the Importer interface to import data from CSV
type CSVImporter struct{}

func (i *CSVImporter) Import(r io.Reader) ([]Customer, error) {
    customers := []Customer{}
    csvReader := csv.NewReader(r)
    for {
        record, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        customers = append(customers, Customer{
            ID:       record[0],
            Name:      record[1],
            Email:     record[2],
            Phone:     record[3],
            Address:   record[4],
        })
    }
    return customers, nil
}

// JSONImporter implements the Importer interface to import data from JSON
type JSONImporter struct{}

func (i *JSONImporter) Import(r io.Reader) ([]Customer, error) {
    customers := []Customer{}
    if err := json.NewDecoder(r).Decode(&customers); err != nil {
        return nil, err
    }
    return customers, nil
}

// XMLImporter implements the Importer interface to import data from XML
type XMLImporter struct{}

func (i *XMLImporter) Import(r io.Reader) ([]Customer, error) {
    customers := []Customer{}
    if err := xml.NewDecoder(r).Decode(&customers); err != nil {
        return nil, err
    }
    return customers, nil
}

// CSVExporter implements the Exporter interface to export data to CSV
type CSVExporter struct{}

func (e *CSVExporter) Export(w io.Writer, customers []Customer) error {
    csvWriter := csv.NewWriter(w)
    defer csvWriter.Flush()
    // Write header
    if err := csvWriter.Write([]string{"ID", "Name", "Email", "Phone", "Address"}); err != nil {
        return err
    }
    // Write customer records
    for _, customer := range customers {
        if err := csvWriter.Write([]string{customer.ID, customer.Name, customer.Email, customer.Phone, customer.Address}); err != nil {
            return err
        }
    }
    return nil
}
 
 // JSONExporter implements the Exporter interface to export data to JSON
type JSONExporter struct{}

func (e *JSONExporter) Export(w io.Writer, customers []Customer) error {
    return json.NewEncoder(w).Encode(customers)
}

// XMLExporter implements the Exporter interface to export data to XML
type XMLExporter struct{}

func (e *XMLExporter) Export(w io.Writer, customers []Customer) error {
    return xml.NewEncoder(w).Encode(customers)