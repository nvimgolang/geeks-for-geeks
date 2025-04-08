package main

import "fmt"

func main() {
	// here rvariable is a array
	rvariable := []string{"GFG", "Geeks", "GeeksforGeeks"}

	// i and j stores the value of rvariable
	// i store index number of individuals string and
	// j store individuals string of the given array
	for i, j := range rvariable {
		fmt.Println(i, j)
	}
}
