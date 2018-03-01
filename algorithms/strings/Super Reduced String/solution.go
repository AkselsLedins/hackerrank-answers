package main

import "fmt"

/* recursive solution */
func superReducedString(s []rune) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return superReducedString(append(s[:i], s[i+2:]...))
		}
	}

	return string(s)
}

func main() {
	var s string
	fmt.Scanf("%v", &s)

	result := superReducedString([]rune(s))

	if len(result) == 0 {
		result = "Empty String"
	}
	fmt.Printf("%s\n", result)
}
