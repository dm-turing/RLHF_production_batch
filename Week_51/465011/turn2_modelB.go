package main

import (
	"fmt"
	"sync"
)

// Employee represents an employee in an organization.
type Employee struct {
	Name       string
	Age        int
	Salary     float64
	Department *Department
}

// String implements the fmt.Stringer interface for Employee.
func (e *Employee) String() string {
	// Lazy computation of Department string representation
	deptStr := func() string {
		if e.Department == nil {
			return "nil"
		}
		return e.Department.String()
	}
	return fmt.Sprintf("Employee{Name: %s, Age: %d, Salary: %g, Department: %s}", e.Name, e.Age, e.Salary, deptStr())
}

// Department represents a department in an organization.
type Department struct {
	Name     string
	Location string
	Manager  *Employee
}

// String implements the fmt.Stringer interface for Department.
func (d *Department) String() string {
	// Lazy computation of Manager's string representation
	mgrStr := func() string {
		if d.Manager == nil {
			return "nil"
		}
		return d.Manager.String()
	}
	return fmt.Sprintf("Department{Name: %s, Location: %s, Manager: %s}", d.Name, d.Location, mgrStr())
}

var once sync.Once

func main() {
	// Create a nested data structure
	manager := &Employee{Name: "Alice", Age: 35, Salary: 75000}
	devDepartment := &Department{Name: "Development", Location: "New York", Manager: manager}
	employee1 := &Employee{Name: "Bob", Age: 28, Salary: 60000, Department: devDepartment}

	// Print the employee details, triggering lazy evaluation of strings
	fmt.Println(employee1)
}
