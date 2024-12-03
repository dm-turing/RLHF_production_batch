package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Example of vector addition using gonum
	a := []float64{1.0, 2.0, 3.0}
	b := []float64{4.0, 5.0, 6.0}
	c := make([]float64, len(a))
	floats.AddTo(c, a, b)
	fmt.Println("Vector addition result:", c)

	// Example of matrix multiplication using gonum
	matA := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	matB := mat.NewDense(2, 2, []float64{5, 6, 7, 8})
	matC := mat.NewDense(2, 2, nil)
	matC.Mul(matA, matB)
	fmt.Printf("Matrix multiplication result:\n%v\n", mat.Formatted(matC))
}
