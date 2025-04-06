package main

import "fmt"

func main() {
	// variables
	str1 := "GeeksforGeeks"
	str2 := "geeksForgeeks"
	str3 := "GeeksforGeeks"
	result1 := str1 == str2
	result2 := str1 == str3

	// display the result
	fmt.Println(result1)
	fmt.Println(result2)

	// display the type of
	// result1 and result2
	fmt.Printf(
		"the type of result1 is %T and "+"the type of result2 is %T",
		result1,
		result2,
	)
}
