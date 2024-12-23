// Go code:
package main

import "fmt"

func main() {
	const pi = 3.14
	radius := 15.5
	volume := fmt.Sprintf("The volume of a sphere with radius %.2f is approximately %.2f.", radius, 4.0/3.0*pi*radius*radius*radius)
	fmt.Println(volume)
}
