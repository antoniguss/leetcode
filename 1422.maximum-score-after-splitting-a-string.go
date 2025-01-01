package main

import "fmt"

// @leet start
func maxScore(s string) (max int) {

	for i := 1; i < len(s); i++ {
		l := s[:i]
		r := s[i:]
		score := countChar(l, '0') + countChar(r, '1')
		if score > max {
			max = score
		}
	}
	return max

}

func countChar(s string, c rune) (count int) {
	for _, c2 := range s {
		if c == c2 {
			count++
		}
	}

	return count
}

// @leet end
func main() {
	fmt.Println(maxScore("011101"))
}

