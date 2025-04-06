package main

import "fmt"

func main() {
	var x float32 = 5.00
	var y float32 = 2.25

	// addition
	fmt.Printf("addition: %g + %g = %g \n", x, y, x+y)

	// substraction
	fmt.Printf("substraction: %g - %g = %g \n", x, y, x-y)

	// multiplication
	fmt.Printf("multiplication: %g * %g = %g \n", x, y, x*y)

	// division
	fmt.Printf("division: %g / %g = %g \n", x, y, x/y)
}
