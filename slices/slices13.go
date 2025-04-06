package main

import "fmt"

func main() {
	// creating slices
	s1 := []int{12, 34, 56}
	var s2 []int

	// checking if the given slice is nil or not
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
}
