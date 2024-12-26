package main

import (
	"fmt"
)

type Matrix [][]float64

func (m Matrix) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		m.formatVerbose(s)
	case 's':
		m.formatString(s)
	default:
		fmt.Fprintf(s, "%%!%c(Matrix=%+v)", verb, m)
	}
}

func (m Matrix) formatVerbose(s fmt.State) {
	rows := len(m)
	if rows == 0 {
		fmt.Fprint(s, "empty matrix")
		return
	}

	cols := len(m[0])

	// Check if all rows have the same number of columns
	for _, row := range m {
		if len(row) != cols {
			fmt.Fprint(s, "irregular matrix")
			return
		}
	}

	if s.Flag('+') {
		fmt.Fprintf(s, "Matrix (%dx%d):\n", rows, cols)
	}

	for _, row := range m {
		for _, val := range row {
			fmt.Fprintf(s, "%.2f ", val)
		}
		fmt.Fprintln(s)
	}
}

func (m Matrix) formatString(s fmt.State) {
	rows := len(m)
	if rows == 0 {
		fmt.Fprint(s, "[]")
		return
	}

	cols := len(m[0])

	// Check if all rows have the same number of columns
	for _, row := range m {
		if len(row) != cols {
			fmt.Fprint(s, "[]")
			return
		}
	}

	fmt.Fprintf(s, "[")
	for i, row := range m {
		fmt.Fprintf(s, "[%.2f ", row[0])
		for j := 1; j < cols; j++ {
			fmt.Fprintf(s, "%.2f ", row[j])
		}
		if i < rows-1 {
			fmt.Fprintf(s, "],\n[")
		} else {
			fmt.Fprintf(s, "]")
		}
	}
	fmt.Fprintf(s, "]")
}

func main() {
	matrix := Matrix{
		{1.2, 3.4, 5.6},
		{7.8, 9.0, 1.2},
		{3.4, 5.6, 7.8},
	}

	fmt.Println(matrix)         // Verbose output
	fmt.Printf("%+v\n", matrix) // Verbose output with flag
	fmt.Printf("%s\n", matrix)  // String representation
	fmt.Printf("%q\n", matrix)  // Quoted string representation
}
