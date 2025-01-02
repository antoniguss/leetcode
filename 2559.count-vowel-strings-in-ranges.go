package main

// @leet start
func vowelStrings(words []string, queries [][]int) []int {
	result := make([]int, len(queries))

	upTo := make([]int, len(words))

	count := 0
	for i, w := range words {
		if isVowel(w[0]) && isVowel(w[len(w)-1]) {
			count++
		}
		upTo[i] = count
	}

	for i, q := range queries {
		if q[0] == 0 {
			result[i] = upTo[q[1]]
		} else {
			result[i] = upTo[q[1]] - upTo[q[0]-1]
		}
	}

	return result
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

// @leet end
func main() {
	vowelStrings([]string{}, [][]int{})
}

