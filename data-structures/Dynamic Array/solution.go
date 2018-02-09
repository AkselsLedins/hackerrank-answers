// NOTE to run it
// go run solution.go < input.txt

package main

import "fmt"

// NOTE Query: 1 x y
//      (1) Find the sequence, seq, at index (( x ^ lastAnswer ) % N) in seqList.
//      (2) Append integer y to sequence .
func query1(seqList [][]int, x, y, lastAnswer int) {
	// (1)
	index := (x ^ lastAnswer) % len(seqList)
	seq := seqList[index]

	// (2)
	seqList[index] = append(seq, y)
}

// NOTE Query: 2 x y
//      (1) Find the sequence, seq, at index (( x ^ lastAnswer ) % N) in seqList.
//      (2) Find the value of element y % size in seq (where size is the size of seq)
//      (3) and assign it to lastAnswer.
//      (4) Print the new value of lastAnswer on a new line
func query2(seqList [][]int, x, y int, lastAnswer *int) {
	// (1)
	index := (x ^ (*lastAnswer)) % len(seqList)
	seq := seqList[index]

	// (2)
	value := seq[y%len(seq)]

	// (3)
	*lastAnswer = value

	// (4)
	fmt.Printf("%d\n", *lastAnswer)
}

func main() {
	// NOTE Input Format: The first line contains two space-separated integers
	//      - N (the number of sequences)
	//      - Q (the number of queries)
	var N, Q int
	fmt.Scanf("%d %d", &N, &Q)

	// NOTE Create a list, seqList, of N empty sequences,
	//      where each sequence is indexed from 0 to N - 1.
	//      The elements within each of the N sequences also use 0-indexing
	seqList := make([][]int, N)

	// NOTE Create an integer lastAnswer and initialize it to 0.
	var lastAnswer int

	for i := 0; i < Q; i++ {
		var q, x, y int
		fmt.Scanf("%d %d %d", &q, &x, &y)
		if q == 1 {
			query1(seqList, x, y, lastAnswer)
		} else if q == 2 {
			query2(seqList, x, y, &lastAnswer)
		}
	}
}
