// @leet start

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	nums := make([]int, length)
	// Merge two arrays
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len(nums1) {
		nums[k] = nums1[i]
		k++
		i++
	}

	for j < len(nums2) {
		nums[k] = nums2[j]
		k++
		j++
	}
	fmt.Println(nums)

	if length%2 == 1 {
		return float64(nums[length/2])
	}

	return float64(nums[length/2]+nums[length/2-1]) / 2.0
}

// @leet end
