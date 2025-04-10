// go program to illustrate the
// concept of main() function

// declaration of the main package
package main

// importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// main function
func main() {
	// sorting the given slice
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice: ", s)

	// finding the index
	fmt.Println("Index value: ", strings.Index("GeeksforGeeks", "ks"))

	// finding the time
	fmt.Println("Time: ", time.Now().Unix())
}
