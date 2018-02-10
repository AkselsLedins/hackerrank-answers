// NOTE to run it
// go run solution.go < input.txt

package main

import (
	"fmt"
	"strings"
)

func main() {
	// NOTE Given an array of n integers and a number, d, perform d left rotations on the array.
	//      Then print the updated array as a single line of space-separated integers.
	var n, d int
	fmt.Scanf("%d %d", &n, &d)

	// Memory Complexity O(n)
	arrayOfNIntegers := make([]int, n)
	// Time Complexity O(n)
	for index := 0; index < n; index++ {
		newPosition := index - d
		if newPosition < 0 {
			newPosition = len(arrayOfNIntegers) + newPosition
		}
		fmt.Scanf("%d", &arrayOfNIntegers[newPosition])
	}

	// NOTE Print a single line of n space-separated integers denoting the final state
	//      of the array after performing d left rotations.
	fmt.Println(strings.Trim(fmt.Sprint(arrayOfNIntegers), "[]"))
}
