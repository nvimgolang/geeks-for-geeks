package main

import "fmt"

// multiply() multiplies two integers and returns the results
func multiply(a, b int) int {
	return a * b
}

func main() {
	result := multiply(5, 10)
	fmt.Printf("multiplication: %d", result)
}
