// NOTE to run it
// go run solution.go < input.txt

package main

import "fmt"

func main() {
	var tripletA, tripletB [3]int
	var aPoints, bPoints int

	fmt.Scanf("%d %d %d", &tripletA[0], &tripletA[1], &tripletA[2])
	fmt.Scanf("%d %d %d", &tripletB[0], &tripletB[1], &tripletB[2])

	for index := 0; index < 3; index++ {
		if tripletA[index] > tripletB[index] {
			aPoints += 1
		} else if tripletB[index] > tripletA[index] {
			bPoints += 1
		}
	}

	fmt.Printf("%d %d", aPoints, bPoints)
}
