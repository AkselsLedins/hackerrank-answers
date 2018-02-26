package main

import (
	"fmt"
	"strconv"
)

func main() {
	// NOTE A single string s containing a time in 12-hour clock format (i.e.: hh:mm:ss:AM or hh:mm:ss:PM)
	var s string
	fmt.Scanf("%s", &s)

	hh, _ := strconv.Atoi(s[0:2])
	mm, _ := strconv.Atoi(s[3:5])
	ss, _ := strconv.Atoi(s[6:8])
	amPm := s[8]

	isPm := false
	if amPm == 'P' {
		isPm = true
	}

	if isPm && hh < 12 {
		hh += 12
	}

	if !isPm && hh == 12 {
		hh = 0
	}
	if isPm && hh < 12 {
		hh += 12
	}

	fmt.Printf("%02d:%02d:%02d\n", hh, mm, ss)
}
