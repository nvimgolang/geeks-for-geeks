package main

import "fmt"

func main() {
	p := 34
	q := 20

	// & (bitwise and)
	result1 := p & q
	fmt.Printf("result of p & q = %d \n", result1)

	// | (bitwise or)
	result2 := p | q
	fmt.Printf("result of p | q = %d \n", result2)

	// ^ (bitwise XOR)
	result3 := p ^ q
	fmt.Printf("result of p ^ q = %d \n", result3)

	// << (left shift)
	result4 := p << 1
	fmt.Printf("result of p << 1 = %d \n", result4)

	// >> (right shift)
	result5 := p >> 1
	fmt.Printf("result of p >> 1 = %d \n", result5)

	// &^ (and not)
	result6 := p &^ q
	fmt.Printf("result of p &^ = %d \n", result6)
}
