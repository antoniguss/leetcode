package main

// @leet start
func removeElement(nums []int, val int) (k int) {

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// @leet end

