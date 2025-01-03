package main

import "fmt"

// @leet start
func waysToSplitArray(nums []int) int {

	sum := 0
	for _, n := range nums {
		sum += n
	}

	count := 0
	curSum := 0
	for i := 0; i < len(nums)-1; i++ {
		curSum += nums[i]

		if curSum >= sum-curSum {
			count++
		}
	}

	return count
}

// @leet end

func main() {
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))

}

