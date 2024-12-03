package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Example of complex number operations
	complex1 := complex(1.0, 2.0) // 1 + 2i
	complex2 := complex(3.0, 4.0) // 3 + 4i

	sum := complex1 + complex2
	product := complex1 * complex2
	modulus := cmplx.Abs(complex1)
	phase := cmplx.Phase(complex1)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
	fmt.Println("Modulus:", modulus)
	fmt.Println("Phase:", phase)
}
