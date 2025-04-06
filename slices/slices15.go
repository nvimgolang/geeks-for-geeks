package main

import (
	"fmt"
	"sort"
)

func main() {
	// creating slice
	slc1 := []string{"python", "java", "c#", "go", "ruby"}
	slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

	fmt.Println("before sorting:")
	fmt.Println("slice 1: ", slc1)
	fmt.Println("slice 2: ", slc2)

	// performing sort operation on the
	// slice using sort function
	sort.Strings(slc1)
	sort.Ints(slc2)

	fmt.Println("\nAfter sorting:")
	fmt.Println("slice 1: ", slc1)
	fmt.Println("slice 2: ", slc2)
}
