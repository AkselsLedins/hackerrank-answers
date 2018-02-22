package main

import "fmt"

func main() {
	var sizeOfArray int
	var sum int

	fmt.Scanf("%d", &sizeOfArray)
	for index := 0; index < sizeOfArray; index++ {
		var n int
		fmt.Scanf("%d", &n)
		sum += n
	}

	fmt.Println(sum)
}
