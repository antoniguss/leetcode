package main

import "fmt"

// @leet start
func myAtoi(s string) (result int) {
	len := len(s)

	i := 0
	//Step 1: Whitspace
	for i < len && s[i] == ' ' {
		i++
	}

	if i == len {
		return 0
	}

	//Step 2: Check sign
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		sign = 1
		i++
	}

	//Step 3: Conversion
	for ; i < len && s[i] >= '0' && s[i] <= '9'; i++ {

		result = result*10 + int(s[i]) - '0'
		if sign*result > 1<<31-1 {
			return 1<<31 - 1
		}

		if sign*result < -1<<31 {
			return -1 << 31
		}
	}

	return sign * result
}

// @leet end

func main() {
	fmt.Println(myAtoi("+"))

}
