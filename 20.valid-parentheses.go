package main

import "fmt"

// @leet start
func isValid(s string) bool {

	stack := make([]rune, 0)

	for _, c := range s {

		switch c {
		case '{', '(', '[':
			stack = append(stack, c)
		case '}':
			if len(stack) == 0 {
				return false
			}
			x := stack[len(stack)-1]
			if x != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 {
				return false
			}
			x := stack[len(stack)-1]
			if x != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 {
				return false
			}
			x := stack[len(stack)-1]
			if x != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0
}

// @leet end
func main() {
	fmt.Println(isValid("()({})"))

}

