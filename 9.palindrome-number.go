// @leet start
package main

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	return x == reverseNum(x)

}

func reverseNum(x int) (result int) {

	result += x % 10
	x /= 10
	for x > 0 {
		result *= 10
		result += x % 10
		x /= 10

	}
	return result

}

// @leet end

