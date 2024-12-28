package main

// @leet start
func twoSum(nums []int, target int) []int {
	indices := make(map[int]int)
	for index1, value := range nums {
		if index2, has := indices[target-value]; has {
			return []int{index1, index2}
		}
		indices[value] = index1

	}
	return []int{-1, -1}

}

// @leet end
