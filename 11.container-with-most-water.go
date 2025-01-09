package main

// @leet start
func maxArea(height []int) (max int) {

	l, r := 0, len(height)-1
	for r > l {

		if height[l] <= height[r] {
			amount := height[l] * (r - l)
			if amount > max {
				max = amount
			}

			l++
		} else {
			amount := height[r] * (r - l)
			if amount > max {
				max = amount
			}

			r--
		}

	}

	return max

}

// @leet end

func main() {

}
