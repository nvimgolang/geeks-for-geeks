package main

import "fmt"

func main() {
	// multiple variables of different types
	// are declared and initialized in the single line
	myvariable1, myvariable2, myvariable3 := 2, "GFG", 67.56

	// display the value and
	// type of the variables
	fmt.Printf("the value of myvariable1 is: %d \n", myvariable1)
	fmt.Printf("the type of the myvariable1 is: %T \n", myvariable1)

	fmt.Printf("the value of myvariable2 is: %s \n", myvariable2)
	fmt.Printf("the type of myvariable2 is: %T \n", myvariable2)

	fmt.Printf("the value of myvariable3 is: %f \n", myvariable3)
	fmt.Printf("the type of myvariable3 is: %T \n", myvariable3)
}
