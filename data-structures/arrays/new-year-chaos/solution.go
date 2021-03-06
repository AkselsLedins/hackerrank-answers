package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int) {
	swaps := 0

	// test for chaos
	for k, v := range q {
		if v-1-k > 2 {
			fmt.Println("Too chaotic")
			return
		}
	}

	// compute swaps
	hasSwapped := true
	for hasSwapped == true {
		hasSwapped = false
		for k := range q {
			if k+1 >= len(q) {
				continue
			}
			if q[k] > q[k+1] {
				// swap !
				q[k], q[k+1] = q[k+1], q[k]
				hasSwapped = true
				swaps++
			}
		}
	}

	fmt.Println(swaps)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
