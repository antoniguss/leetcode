package main

import "fmt"

// @leet start
func wordSubsets(words1 []string, words2 []string) []string {

	var maxCounts [26]int

	for _, word := range words2 {

		occ := occurances(word)

		for char, freq := range occ {
			if maxCounts[char] < freq {
				maxCounts[char] = freq
			}
		}
	}

	result := make([]string, 0)
	for _, word := range words1 {

		occ := occurances(word)

		isUniversal := true

		for i := range 26 {
			if occ[i] < maxCounts[i] {
				isUniversal = false
				break
			}

		}

		if isUniversal {
			result = append(result, word)
		}

	}

	return result

}

func occurances(word string) []int {
	result := make([]int, 26)

	for _, c := range word {
		result[c-'a']++
	}
	return result
}

// @leet end
func main() {
	fmt.Println(occurances("apple"))

}
