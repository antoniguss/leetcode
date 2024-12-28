// @leet start
func lengthOfLongestSubstring(s string) int {
	var maxLen int
	used := make(map[byte]bool)
	for i, j := 0, 0; i < len(s); i++ {
		for j < len(s) && !used[s[j]] {
			used[s[j]] = true
			j++
			if j-i > maxLen {
				maxLen = j - i
			}

		}

		used[s[i]] = false

	}

	return maxLen

}

// @leet end
