package main

import "fmt"

func main() {
	var day interface{} = 4
	switch v := day.(type) {
	case int:
		switch v {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		default:
			fmt.Println("Invalid day")

		}
	default:
		fmt.Printf("Unknown type: $T \n", v)
	}
}
