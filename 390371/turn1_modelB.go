package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// ModularExponentiation calculates a^b mod m with adaptive 'bit precision' checks
func ModularExponentiation(a *big.Int, b *big.Int, m *big.Int) *big.Int {
	result := big.NewInt(1)
	base := new(big.Int).Set(a)

	for b.BitLen() > 0 {
		// Check if the least significant bit is set
		if b.Bit(0) == 1 {
			result = new(big.Int).Mul(result, base).Mod(a, b)
		}
		// Square the base and reduce modulo m
		base = new(big.Int).Mul(base, base).Mod(a, b)
		// Right shift b by one
		b = new(big.Int).Rsh(b, 1)
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter base (a): ")
	scanner.Scan()
	a, _ := new(big.Int).SetString(scanner.Text(), 10)

	fmt.Print("Enter exponent (b): ")
	scanner.Scan()
	b, _ := new(big.Int).SetString(scanner.Text(), 10)

	fmt.Print("Enter modulus (m): ")
	scanner.Scan()
	m, _ := new(big.Int).SetString(scanner.Text(), 10)

	result := ModularExponentiation(a, b, m)
	fmt.Printf("%d^%d mod %d = %v\n", a, b, m, result)
}
