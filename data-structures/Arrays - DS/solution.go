// NOTE to run it
// go run solution.go < input.txt

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	// we create an array of ints of size n
	arrayOfInts := make([]int, n)

	// we fill the array
	for i := n - 1; i >= 0; i-- {
		fmt.Scanf("%d", &arrayOfInts[i])
	}

	// we print the array in reverse
	for _, e := range arrayOfInts {
		fmt.Printf("%d ", e)
	}

	// one-liner alternative to print the array but its harder to understand
	// fmt.Println(strings.Trim(fmt.Sprint(arrayOfInts), "[]"))
}
