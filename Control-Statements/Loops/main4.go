package main

import "fmt"

func main() {
	// string as a range in the for loop
	for i, j := range "XabCd" {
		fmt.Printf("The index number of %U is %d \n", j, i)
	}
}
