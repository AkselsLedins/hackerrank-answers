package main

import "fmt"

type Kangaroo struct {
	Position      int
	MetersPerJump int
}

func canTheyMeet(k1, k2 Kangaroo) string {
	// this is equal to solve this equation
	//    k1.Position + k1.MetersPerJump * x = k2.Position + k2.MetersPerJump * x
	// |_ k1.Position - k2.Position = k2.MetersPerJump * x - k1.MetersPerJump * x
	// |_ k1.Position - k2.Position = x * (k2.MetersPerJump - k1.MetersPerJump)
	// |       k1.Position - k2.Position          =    x * (k2.MetersPerJump - k1.MetersPerJump)
	// |_         -------------------------           -----------------------------------------
	//      (k2.MetersPerJump - k1.MetersPerJump)       (k2.MetersPerJump - k1.MetersPerJump)
	//
	// We can't make a half jump, this means:
	// (k1.Position - k2.Position) % (k2.MetersPerJump - k1.MetersPerJump) == 0,
	// if the two kangaroos will met someday

	if (k1.Position-k2.Position)%(k2.MetersPerJump-k1.MetersPerJump) == 0 {
		return "YES"
	}

	return "NO"
}

func main() {
	// NOTE input format
	// 4 space-separated integers
	// x1, x2, x3, x4 : starting locations xRoo and meters per jump vRoo for kangaroos 1 and 2

	var kangaroo1, kangaroo2 Kangaroo
	fmt.Scanf("%d %d %d %d",
		&kangaroo1.Position,
		&kangaroo1.MetersPerJump,
		&kangaroo2.Position,
		&kangaroo2.MetersPerJump,
	)

	if kangaroo2.MetersPerJump >= kangaroo1.MetersPerJump && kangaroo1.Position != kangaroo2.Position {
		fmt.Printf("NO\n")
		return
	}

	response := canTheyMeet(kangaroo1, kangaroo2)

	fmt.Printf("%v\n", response)
}
