package main

import "fmt"

// 0 = safe
// 1 = avoid

func main() {
	var numberOfClouds int

	fmt.Scanf("%d", &numberOfClouds)

	clouds := make([]int, numberOfClouds)

	for i := 0; i < numberOfClouds; i++ {
		fmt.Scanf("%d", &clouds[i])
	}

	jumps := 0

	for i := 0; i < numberOfClouds-1; {
		jumps++
		// test big jump (+2)
		if i+2 < numberOfClouds {
			if clouds[i+2] == 0 {
				i += 2
				continue
			}
		}
		// small jump
		i++
	}

	fmt.Println(jumps)
}
