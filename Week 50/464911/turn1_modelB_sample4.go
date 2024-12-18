package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Salary float64
}

func main() {
	employee := Employee{Person{"Alice", 30}, 50000}
	fmt.Println(employee.Name) // Output: Alice
}
