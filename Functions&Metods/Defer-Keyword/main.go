package main

import "fmt"

func mul(a1, a2 int) int {
	res := a1 * 2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello!, GeeksforGeeks")
}

func main() {
	// calling mul() function
	// here mul function behaves
	// like a normal function
	mul(23, 45)

	// calling mul() function
	// using defer keyword
	// here the mul() function
	// is defer function
	defer mul(23, 56)

	// calling show() fucntion
	show()
}
