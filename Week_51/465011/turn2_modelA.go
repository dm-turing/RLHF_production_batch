package main

import (
	"fmt"
)

// Variable represents a mathematical variable and supports lazy string evaluation.
type Variable struct {
	name string
}

// String implements the Stringer interface and returns a lazy string.
func (v *Variable) String() string {
	return fmt.Sprintf("Computing variable name: %s", v.name)
}

// Term represents a mathematical term (constant or variable) with lazy string evaluation.
type Term struct {
	val  float64
	_var *Variable
}

// String implements the Stringer interface for the Term.
func (t *Term) String() string {
	return fmt.Sprintf("%.2f%s", t.val, t._var.String())
}

// Equation represents a mathematical equation with lazy string evaluation.
type Equation struct {
	left  *Term
	right *Term
}

// String implements the Stringer interface for the Equation.
func (e *Equation) String() string {
	return fmt.Sprintf("%.2f %s = %.2f %s", e.left.val, e.left._var.String(), e.right.val, e.right._var.String())
}

func main() {
	// Create variable, term, and equation instances with lazy evaluation
	x := &Variable{name: "x"}
	leftTerm := &Term{val: 2.5, _var: x}
	rightTerm := &Term{val: 3.1, _var: x}
	equation := &Equation{left: leftTerm, right: rightTerm}

	// String representation is computed only when printed
	fmt.Println("Equation:", equation)
}
