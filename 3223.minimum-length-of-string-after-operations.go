package main

// @leet start
func minimumLength(s string) int {

	counts := char_counts(s)

	result := 0
	for _, count := range counts {
		if count >= 3 {
			result += 2 - count%2
		} else {
			result += count
		}
	}

	return result

}

func char_counts(word string) map[rune]int {
	result := make(map[rune]int)

	for _, c := range word {
		result[c-'a']++
	}
	return result
}

// @leet end

