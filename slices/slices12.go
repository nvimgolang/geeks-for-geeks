package main

import "fmt"

func main() {
	// creating a zero value slice
	arr := []int{55, 66, 77, 88, 99, 22}
	slc := arr[0:4]

	// before modifying

	fmt.Println("original_array: ", arr)
	fmt.Println("original_slice: ", slc)

	// after modification
	slc[0] = 100
	slc[1] = 1000
	slc[2] = 1000

	fmt.Println("\nNew_array: ", arr)
	fmt.Println("New_Slice: ", slc)
}
