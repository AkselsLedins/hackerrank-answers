package main

import "fmt"

func main() {
	var arraySize int
	fmt.Scanf("%d", &arraySize)

	numberOfMatchingSocks := 0
	hashtable := make(map[int]bool)

	for i := 0; i < arraySize; i++ {
		var color int
		fmt.Scanf("%d", &color)

		if hashtable[color] == true {
			hashtable[color] = false
			numberOfMatchingSocks++
		} else {
			hashtable[color] = true
		}
	}

	fmt.Println(numberOfMatchingSocks)
}
