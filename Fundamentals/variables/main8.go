package main

import "fmt"

func main() {
	myvar1, myvar2, myvar3 := 800, "GeeksforGeeks", 47.56

	fmt.Printf("the value of myvar1 is: %d \n", myvar1)
	fmt.Printf("the type of myvar1 is: %T \n", myvar1)

	fmt.Printf("the value of myvar2 is: %s \n", myvar2)
	fmt.Printf("the type of myvar2 is: %T \n", myvar2)

	fmt.Printf("the value of myvar3 is: %f \n", myvar3)
	fmt.Printf("the type of myvar3 is: %T \n", myvar3)
}
