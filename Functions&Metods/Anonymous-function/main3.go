pckage main

import "fmt"

func main() {
	// passing arguments in anonymous function
	func(ele string) {
		fmt.Println(ele)
	}("GeekforGeeks)

