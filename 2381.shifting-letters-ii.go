package main

import (
	"fmt"
	"strings"
)

// @leet start
func shiftingLetters(s string, shifts [][]int) string {
	//Cumulative shifts
	cShifts := make([]int, len(s)+1)
	for _, shift := range shifts {

		if shift[2] == 1 {
			cShifts[shift[0]]++
			cShifts[shift[1]+1]--
		} else {
			cShifts[shift[0]]--
			cShifts[shift[1]+1]++
		}
	}

	// Apply final calculated shift
	result := []byte(s)
	shift := 0
	for i, c := range s {
		shift += cShifts[i]

		newChar := 'a' + ((int(c)-'a'+shift)%26+26)%26
		result[i] = byte(newChar)
	}

	return string(result)
}

// @leet end

func main() {
	fmt.Println(shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}))
}

