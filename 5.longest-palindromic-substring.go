package main

// @leet start
func longestPalindrome(s string) (result string) {

	var left, right, length int
	for i := range len(s) {

		// Case 1: grow palindrome from character at i, palindrome will have an
		// (odd length palindrome)

		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > length {
			length = r - l - 1
			left = l + 1
			right = r
		}

		// Case 2: grow palindrome from character i and i+1 (assuming both are the same)
		// (even length palindrome)

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}

		if r-l-1 > length {
			length = r - l - 1
			left = l + 1
			right = r
		}

	}

	return s[left:right]

}

// @leet end
func main() {
}
