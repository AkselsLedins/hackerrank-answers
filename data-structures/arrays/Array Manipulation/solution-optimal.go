// NOTE to run it
// go run solution-optimal.go < input.txt

/* ************************************************************************** */
/* I take no credit for this solution, I only implemented amansbhandar one    */
/* from hackerrank in GoLang                                                  */
/* ************************************************************************** */

/* ************************************************************************** */
/* real	0m1.520s
/* user	0m0.826s                    It's fast ♪~ ᕕ(ᐛ)ᕗ
/* sys	0m0.767s
/*
/* vs solution.go implementation
/*
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
    slice := make([]int64, n+1)

    var index int64
    for index = 0; index < m; index++ {
        // NOTE Next m lines will contain three integers a, b and k separated
        //      by a single space
        var a, b, k int64
        fmt.Scanf("%d %d %d", &a, &b, &k)
        slice[a] += k
        if b+1 <= n {
            slice[b+1] -= k
        }
    }

    var max int64
    var x int64
    max = math.MinInt64
    for index = 1; index <= n; index++ {
        x += slice[index]
        if max < x {
            max = x
        }
    }

    // NOTE Print in a single line the maximum value in the updated slice.
    fmt.Println(max)
}
