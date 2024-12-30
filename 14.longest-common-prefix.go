package main

import "fmt"

// @leet start
func longestCommonPrefix(strs []string) string {

	l := 0
	matches := true
	for ; matches && l < len(strs[0]); l++ {

		for i := 1; i < len(strs); i++ {

			if l == len(strs[i]) || strs[i][l] != strs[0][l] {
				matches = false
				break
			}
		}

		if !matches {
			return strs[0][:l]
		}
	}

	return strs[0]

}

// @leet end

func main() {
	fmt.Print(longestCommonPrefix([]string{"flower", "flow"}))
}

