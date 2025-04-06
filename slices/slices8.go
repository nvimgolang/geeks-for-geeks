package main

import "fmt"

func main() {
	// creating a slcie
	myslice := []string{"This", "is", "the", "tutorial", "of", "Go", "language"}

	// iterate using for loop
	for e := 0; e < len(myslice); e++ {
		fmt.Println(myslice[e])
	}
}
