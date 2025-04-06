package main

import "fmt"

func main() {
	// using short variable declaration
	// multiple variables of same types
	// are declared and initalized in
	// the single line
	myvar1, myvar2, myvar3 := 800, 34, 56

	// display the value and
	// type of the variables
	fmt.Printf("the value of myvar1 is: %d \n", myvar1)
	fmt.Printf("the type of myvar1 is: %T \n", myvar1)

	fmt.Printf("the value of myvar2 is: %d \n", myvar2)
	fmt.Printf("the type of myvar2 is: %T \n", myvar2)

	fmt.Printf("the value of myvar3 is: %d \n", myvar3)
	fmt.Printf("the type of myvar3 is: %T \n", myvar3)
}
