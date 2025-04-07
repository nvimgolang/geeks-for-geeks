package main

import "fmt"

func main() {
	a := 4

	// using address of operator(&) and
	// pointer indirection(*) operator
	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)
}
