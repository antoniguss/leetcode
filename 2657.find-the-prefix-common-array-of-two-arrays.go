package main

// @leet start
func findThePrefixCommonArray(A []int, B []int) []int {

	result := make([]int, len(A))

	frequencies := make([]int, len(A)+1)
	count := 0

	for i := 0; i < len(A); i++ {

		frequencies[A[i]]++

		if frequencies[A[i]] == 2 {
			count++
		}

		frequencies[B[i]]++

		if frequencies[B[i]] == 2 {
			count++
		}

		result[i] = count

	}

	return result

}

// @leet end

