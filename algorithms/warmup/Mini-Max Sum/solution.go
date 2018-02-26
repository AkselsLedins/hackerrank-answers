package main

import (
	"fmt"
	"sort"
)

func main() {
	arrayOf5Ints := make([]int, 5)

	fmt.Scanf("%d %d %d %d %d",
		&arrayOf5Ints[0],
		&arrayOf5Ints[1],
		&arrayOf5Ints[2],
		&arrayOf5Ints[3],
		&arrayOf5Ints[4])

	sum := arrayOf5Ints[0] + arrayOf5Ints[1] + arrayOf5Ints[2] + arrayOf5Ints[3] + arrayOf5Ints[4]

	sort.Ints(arrayOf5Ints)

	fmt.Printf("%d %d\n", sum-arrayOf5Ints[4], sum-arrayOf5Ints[0])
}
