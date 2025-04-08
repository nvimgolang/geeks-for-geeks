package main

import "fmt"

func multiply(a, b *int) int {
	*a = *a * 2 // modifying a's value at its memory address
	return *a * *b
}

func main() {
	x := 5
	y := 10
	fmt.Printf("Before: x = %d, y = %d \n", x, y)
	result := multiply(&x, &y)
	fmt.Printf("multiplication: %d \n", result)
	fmt.Printf("After: x = %d, y = %d \n", x, y)
}
