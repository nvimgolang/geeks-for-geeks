package main

import "fmt"

func main() {
	// variables declared and
	// initalized without the
	// explicit type
	myvariable1 := 20
	myvariable2 := "GeeksforGeeks"
	myvariable3 := 34.80

	// display the value and the
	// type of the variables
	fmt.Printf("the value of myvariable1 is: %d \n", myvariable1)
	fmt.Printf("the type of myvariable1 is: %T \n", myvariable1)

	fmt.Printf("the value of myvariable2 is: %s \n", myvariable2)
	fmt.Printf("the type of myvariable2 is: %T \n", myvariable2)

	fmt.Printf("the value of myvariable3 is: %f \n", myvariable3)
	fmt.Printf("the type of myvariable3 is: %T \n", myvariable3)
}
