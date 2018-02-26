package main

import "fmt"

func main() {
	var nieceAge int

	max := -1

	fmt.Scanf("%d", &nieceAge)

	var candle, numberOfCandles int
	for n := 0; n < nieceAge; n++ {
		fmt.Scanf("%d", &candle)
		if candle > max {
			max = candle
			numberOfCandles = 0
		}
		if candle == max {
			numberOfCandles++
		}
	}
	fmt.Printf("%v\n", numberOfCandles)
}
