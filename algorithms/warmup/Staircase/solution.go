package main

import "fmt"

func main() {
	var stairSize int
	fmt.Scanf("%d", &stairSize)

	for i := 0; i < stairSize; i++ {

		for n := 1; n < stairSize-i; n++ {
			fmt.Printf(" ")
		}
		for n := i + 1; n > 0; n-- {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
