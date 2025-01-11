package main

// @leet start
func canConstruct(s string, k int) bool {

	if len(s) < k {
		return false
	}

	odd_counts := 0
	counts := make([]int, 26)
	for _, c := range s {
		counts[c-'a']++
	}

	for _, v := range counts {
		if v%2 == 1 {
			odd_counts++
		}
	}

	return odd_counts <= k
}

func char_counts(word string) map[rune]int {
	result := make(map[rune]int)

	for _, c := range word {
		result[c-'a']++
	}
	return result
}

// @leet end
