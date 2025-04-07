package main

import "fmt"

func main() {
	var p int = 23
	var q int = 60

	if p != q && p <= q {
		fmt.Println("true")
	}

	if p != q || p <= q {
		fmt.Println("true")
	}

	if !(p == q) {
		fmt.Println("true")
	}
}
