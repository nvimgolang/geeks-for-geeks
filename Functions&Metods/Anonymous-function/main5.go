package main

import "fmt"

// returning anonymous function
func GFG() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "GeekforGeeks"
	}
	return myf
}

func main() {
	value := GFG()
	fmt.Println(value("Welcome ", "to "))
}
