// NOTE to run it
// go run solution.go < input.txt

/* ************************************************************************** */
/* real	8m40.245s
/* user	8m39.444s                   It's slow (ง ͠° ͟ل͜ ͡°)ง
/* sys	0m1.576s
/* ************************************************************************** */

package main

import (
    "fmt"
    "math"
)

func main() {
    // NOTE The first line will contain two integers n and m separated by
    //      a single space
    var n, m int64
    fmt.Scanf("%d %d", &n, &m)

    // NOTE You are given a list(1-indexed) of size , initialized with zeroes
    list := make([]int64, n)

    var max int64
    max = math.MinInt64

    var index int64
    for index = 0; index < m; index++ {
        // NOTE Next m lines will contain three integers a, b and k separated
        //      by a single space
        var a, b, k int64
        fmt.Scanf("%d %d %d", &a, &b, &k)

        for subarrayStart := a - 1; subarrayStart <= b-1; subarrayStart++ {
            list[subarrayStart] += k
            if list[subarrayStart] > max {
                max = list[subarrayStart]
            }
        }
    }

    // NOTE Print in a single line the maximum value in the updated list.
    fmt.Println(max)
}
