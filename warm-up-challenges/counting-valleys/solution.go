package main

import "fmt"

func main() {
	var numberOfSteps int

	fmt.Scanf("%d\n", &numberOfSteps)

	elevation := 0 // sea level
	numberOfValleys := 0

	for i := 0; i < numberOfSteps; i++ {
		var step byte

		fmt.Scanf("%1c", &step)

		if step == 'U' {
			elevation++
			continue
		}

		if step == 'D' {
			if elevation == 0 {
				numberOfValleys++
			}
			elevation--
		}

	}
	fmt.Printf("%d\n", numberOfValleys)
}
