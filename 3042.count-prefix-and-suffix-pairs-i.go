package main

import (
	"fmt"
)

// @leet start
func countPrefixSuffixPairs(words []string) (result int) {

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefix(words[i], words[j]) && isSuffix(words[i], words[j]) {
				result++
			}
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

func isSuffix(suffix, word string) bool {

	if len(word) < len(suffix) {
		return false
	}

	for i := 0; i < len(suffix); i++ {
		if word[len(word)-len(suffix)+i] != suffix[i] {
			return false
		}
	}
	return true
}

// @leet end
func main() {
	fmt.Println(isSuffix("aba", "ababa"))
}

