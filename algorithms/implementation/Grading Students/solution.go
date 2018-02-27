package main

import "fmt"

func main() {
	var arraySize int

	fmt.Scanf("%d", &arraySize)

	var grade int
	for i := 0; i < arraySize; i++ {
		fmt.Scanf("%d\n", &grade)
		if grade < 38 {
			fmt.Printf("%d\n", grade)
			continue
		}

		to5 := 5 - grade%5
		if to5 < 3 {
			grade += to5
		}

		fmt.Printf("%d\n", grade)
	}
}
