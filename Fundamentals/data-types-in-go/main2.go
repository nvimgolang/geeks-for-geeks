package main

import "fmt"

func main() {
	var x int16 = 170
	var y int16 = 83

	// addition
	fmt.Printf("addition: %d + %d = %d \n", x, y, x+y)

	// substraction
	fmt.Printf("substraction: %d - %d = %d \n", x, y, x-y)

	// multiplication
	fmt.Printf("multiplication: %d * %d = %d \n", x, y, x*y)

	// division
	fmt.Printf("division: %d / %d = %d \n", x, y, x/y)

	// modulus
	fmt.Printf("modulus: %d %% %d = %d \n", x, y, x%y)
}
