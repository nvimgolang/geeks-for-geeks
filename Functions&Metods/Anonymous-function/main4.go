package main

import "fmt"

// passing anonymous function as an argument
func GFG(i func(p, q string) string) {
	fmt.Println(i("Geeks", "for"))
}

func main() {
	value := func(p, q string) string {
		return p + q + "Geeks"
	}
	GFG(value)
}
