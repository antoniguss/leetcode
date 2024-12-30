package main

// @leet start
func reverse(x int) (result int) {
	sign := 1
	if x < 0 {
		sign = -1
		x *= -1
	}

	result += x % 10
	x /= 10
	INT_MAX := 1 << 31 / 10

	for x > 0 {
		// If next step will cause overflow
		if result > INT_MAX {
			return 0
		}
		result = result*10 + x%10
		x /= 10

	}
	return result * sign

}

// @leet end
