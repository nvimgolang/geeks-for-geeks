package main

import "fmt"

func main() {
	const A = "GFG"
	B := "GeeksforGeeks"

	// concrat strings
	helloWorld := A + " " + B
	helloWorld += "!"
	fmt.Println(helloWorld)

	// compare strings
	fmt.Println(A == "GFG")
	fmt.Println(B < A)
}
