package main

// @leet start
func prefixCount(words []string, pref string) (result int) {

	for _, word := range words {
		if isPrefix(pref, word) {
			result++
		}
	}

	return result
}

func isPrefix(prefix, word string) bool {

	if len(word) < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if word[i] != prefix[i] {
			return false
		}
	}
	return true

}

// @leet end
func main() {

}

