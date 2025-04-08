package main

import "fmt"

// variadic function to calcualate sum
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// function with a regular parameter and a variable parameter
func greet(prefix string, nums ...int) {
	fmt.Println(prefix)
	for _, n := range nums {
		fmt.Println("Number:", n)
	}
}

func main() {
	greet("Sum of numbers:", 1, 2, 3)
	greet("Sum of numbers:", 4, 5)
	greet("No numbers su)
}
