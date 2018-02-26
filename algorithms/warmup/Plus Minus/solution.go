// NOTE test: go run solution.go < input.txt

package main

import "fmt"

func main() {
	var arraySize int

	var positive, negative, zeroes int

	fmt.Scanf("%d", &arraySize)

	for i := 0; i < arraySize; i++ {
		var number int
		fmt.Scanf("%d", &number)

		if number > 0 {
			positive++
		} else if number < 0 {
			negative++
		} else {
			zeroes++
		}
	}

	fmt.Printf("%f\n", float64(positive)/float64(arraySize))
	fmt.Printf("%f\n", float64(negative)/float64(arraySize))
	fmt.Printf("%f\n", float64(zeroes)/float64(arraySize))
}
