package main

import (
	"fmt"
	"sort"
	"strings"
)

// @leet start
func stringMatching(words []string) []string {

	result := make([]string, 0)

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	for i := 0; i < len(words); i++ {

		for j := i - 1; j >= 0; j-- {

			if len(words[j]) == len(words[i]) {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}

		}

	}

	return result

}

// @leet end

func main() {
	fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))

}
