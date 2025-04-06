package main

import "fmt"

func main() {
	// creating slice
	myslice := []string{"this", "is", "the", "tutorial", "of", "Go", "language"}

	// iterate slice
	// using range in for loop
	for index, ele := range myslice {
		fmt.Printf("Index = %d and element = %s \n", index+3, ele)
	}
}
