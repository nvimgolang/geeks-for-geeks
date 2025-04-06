package main

import "fmt"

func main() {
	// using unsigned 8-bit int
	var x uint8 = 225
	fmt.Println(x, x-3)

	// using 16-bit signed int
	var y int16 = 32767
	fmt.Println(y+2, y-2)
}
