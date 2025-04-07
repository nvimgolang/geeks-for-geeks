package main

import "fmt"

func main() {
	p := 34
	q := 20

	// addition
	result1 := p + q
	fmt.Printf("result of p + q = %d \n", result1)

	// subiteraction
	result2 := p - q
	fmt.Printf("result of p - q = %d \n", result2)

	// multiplication
	result3 := p * q
	fmt.Printf("result of p * q = %d \n", result3)

	// division
	result4 := p / q
	fmt.Printf("result of p / q = %d \n", result4)

	// modulus
	result5 := p % q
	fmt.Printf("result of p %% q = %d \n", result5)
}
