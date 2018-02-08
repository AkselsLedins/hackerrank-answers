// NOTE to run it
// go run solution.go < input.txt

package main

import (
	"fmt"
	"math"
)

func main() {
	var array6x6 [6][6]int

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scanf("%d", &array6x6[i][j])
		}
	}

	max := math.MinInt32

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			sumHourglass := computeSumHourglass(array6x6, i, j)
			if sumHourglass > max {
				max = sumHourglass
			}
		}
	}

	fmt.Printf("%d", max)
}

func computeSumHourglass(arr [6][6]int, i, j int) (sum int) {
	return arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] +
		arr[i][j] +
		arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
}
