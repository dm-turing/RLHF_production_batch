package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Project struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
}

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Department Department `json:"department"`
	Projects   []Project  `json:"projects"`
}

type Company struct {
	Employees []Employee `json:"employees"`
}

func main() {
	jsonData := []byte(`{
		"employees": [
			{
				"id": 1,
				"name": "Alice",
				"department": {
					"id": 101,
					"name": "Engineering"
				},
				"projects": [
					{
						"id": 201,
						"name": "Project Alpha",
						"startDate": "2023-01-15"
					},
					{
						"id": 202,
						"name": "Project Beta",
						"startDate": "2023-03-01"
					}
				]
			},
			{
				"id": 2,
				"name": "Bob",
				"department": {
					"id": 102,
					"name": "Marketing"
				},
				"projects": []
			}
		]
	}`)

	var company Company
	err := json.Unmarshal(jsonData, &company)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	fmt.Println(company)
}
