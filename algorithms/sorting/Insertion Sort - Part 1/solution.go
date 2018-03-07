// go run solution.go < input.txt

package main

import "fmt"

func main() {
	var size int

	// size - the size of the array
	fmt.Scanf("%d", &size)
	// arr the array containing size - 1 sorted integers and 1 unsorted integer e in the rightmost cell
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	// Exceptions:
	if size == 1 {
		fmt.Printf("%d", arr[0])
	}

	needToBreak := false

	lastElement := arr[size-1]
	for i := 2; i <= size; i++ {
		elementToCompare := arr[size-i]
		if elementToCompare > lastElement {
			arr[size-i+1] = arr[size-i]
		} else {
			arr[size-i+1] = lastElement
			needToBreak = true
		}

		// print array
		for _, item := range arr {
			fmt.Printf("%d ", item)
		}
		fmt.Printf("\n")
		if needToBreak {
			break
		}
	}

	// first item (worst case)
	if !needToBreak {
		arr[0] = lastElement
		// print array
		for _, item := range arr {
			fmt.Printf("%d ", item)
		}
		fmt.Printf("\n")
	}
}
