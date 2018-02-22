package main

import "fmt"

func main() {
	var n int
	var aVeryBigSum int64

	fmt.Scanf("%d", &n)

	for index := 0; index < n; index++ {
		var toAdd int64
		fmt.Scanf("%v", &toAdd)
		aVeryBigSum += toAdd
	}

	fmt.Println(aVeryBigSum)
}
