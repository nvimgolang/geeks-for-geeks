package main

import "fmt"

func main() {
	// creating a zero value slice
	var myslice []string
	fmt.Printf("lenght = %d\n", len(myslice))
	fmt.Printf("capacity = %d ", cap(myslice))
}
