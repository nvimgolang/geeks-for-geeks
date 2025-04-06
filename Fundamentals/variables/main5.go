package main

import "fmt"

func main() {
	// using short variable declaration
	myvar1 := 39
	myvar2 := "GeeksforGeeks"
	myvar3 := 34.67

	// display the value and type of the variables
	fmt.Printf(
		"the value of myvar1 is: %d \n"+"the type of myvar1 is: %T \n",
		myvar1, myvar1)

	fmt.Printf(
		"the value of myvar2 is: %s \n"+"the type of myvar2 is: %T \n",
		myvar2,
		myvar2,
	)

	fmt.Printf(
		"the value of myvar3 is: %f \n"+"the type of myvar3 is: %T \n",
		myvar3,
		myvar3,
	)
}
