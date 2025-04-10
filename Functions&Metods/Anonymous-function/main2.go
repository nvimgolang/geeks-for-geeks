package main

import "fmt"

func main() {
	// assigning and anonymous function to a variable
	value := func() {
		fmt.Println("Welcome! to GeekforGeeks")
	}
	value()
}
