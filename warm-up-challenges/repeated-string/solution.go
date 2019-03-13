package main

import "fmt"

func main() {
	var s string
	var n int

	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%d\n", &n)

	aCounter := 0

	for i := 0; i < len(s) && i < n; i++ {
		if s[i] == 'a' {
			aCounter++
		}
	}

	if len(s) > n {
		fmt.Println(aCounter)
		return
	}

	aCounter = aCounter * int(n/len(s))
	remainingLetters := n % len(s)

	for i := 0; i < remainingLetters; i++ {
		if s[i] == 'a' {
			aCounter++
		}
	}

	fmt.Println(aCounter)
}
