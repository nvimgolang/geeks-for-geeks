package main

import "fmt"

func main() {
	const trueConst = true

	// type definition using type keyword
	type myBool bool
	defaultBool := trueConst          // allowed
	var customBool myBool = trueConst // allowed

	// defaultBool = customBool // not allowed
	fmt.Println(defaultBool)
	fmt.Println(customBool)
}
