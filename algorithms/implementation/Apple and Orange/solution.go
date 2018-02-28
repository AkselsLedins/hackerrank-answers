package main

import "fmt"

func addOneIfInsideHoue(count *int, s, t, treePosition int, distance int) {
	if treePosition+distance >= s && treePosition+distance <= t {
		*count++
	}
}

func main() {
	var applesInHouse, orangesInHouse int

	// NOTE s and t, left and right sides of Sam’s house
	var s, t int
	fmt.Scanf("%d %d", &s, &t)

	// NOTE a and b, Larry’s and Rob’s positions in the trees
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// NOTE m and n, number of apples and oranges thrown
	var m, n int
	fmt.Scanf("%d %d", &m, &n)

	var object int

	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &object)
		addOneIfInsideHoue(&applesInHouse, s, t, a, object)
	}

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &object)
		addOneIfInsideHoue(&orangesInHouse, s, t, b, object)
	}

	fmt.Printf("%d\n%d\n",
		applesInHouse,
		orangesInHouse,
	)
}
