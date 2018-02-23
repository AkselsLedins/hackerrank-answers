package main

import (
	"fmt"
	"math"
)

func main() {
	var matrixSize int
	fmt.Scanf("%d", &matrixSize)

	// create matrix
	matrix := make([][]int, matrixSize)
	for i := 0; i < matrixSize; i++ {
		matrix[i] = make([]int, matrixSize)
		for y := 0; y < matrixSize; y++ {
			fmt.Scanf("%d", &matrix[i][y])
		}
	}

	// get diagonals sums
	var primaryDiagoanalSum int
	var secondaryDiagonalSum int
	for i := 0; i < matrixSize; i++ {
		primaryDiagoanalSum += matrix[i][i]
		secondaryDiagonalSum += matrix[matrixSize-i-1][i]
	}

	difference := primaryDiagoanalSum - secondaryDiagonalSum

	fmt.Printf("%v\n", math.Abs(float64(difference)))
}
