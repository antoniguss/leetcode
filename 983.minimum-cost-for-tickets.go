package main

import (
	"fmt"
	"math"
)

// @leet start
func mincostTickets(days []int, costs []int) int {
	maxDay := days[len(days)-1]
	dp := make([]int, maxDay+1)

	travelDays := make([]bool, maxDay+1)
	for _, day := range days {
		travelDays[day] = true
	}

	for d := 1; d < len(dp); d++ {

		if !travelDays[d] {
			dp[d] = dp[d-1]
			continue
		}

		var (
			min0, min1, min2 int
		)
		min0 = costs[0] + dp[d-1]

		if d-7 >= 0 {
			min1 = costs[1] + dp[d-7]
		} else {
			min1 = costs[1]
		}

		if d-30 >= 0 {
			min2 = costs[2] + dp[d-30]
		} else {
			min2 = costs[2]
		}

		dp[d] = min(min0, min1, min2)

	}

	fmt.Println(dp)
	return dp[maxDay]

	// Recursive approach, very inefficient for long inputs
	// Return the minimum of buying a 1-day, a 7-day and a 30-pass
	// if len(days) == 0 {
	// 	return 0
	// }
	//
	// //Buy 1 day pass
	// min0 := costs[0] + mincostTickets(days[1:], costs)
	//
	// //Buy 7 day pass
	// startDay := days[0]
	// endIndex := 0
	// for {
	// 	if endIndex == len(days)-1 || days[endIndex+1]-startDay >= 7 {
	// 		break
	// 	}
	// 	endIndex++
	// }
	// min1 := costs[1] + mincostTickets(days[endIndex+1:], costs)
	// //Buy 30 day pass
	// startDay = days[0]
	// endIndex = 0
	// for {
	// 	if endIndex == len(days)-1 || days[endIndex+1]-startDay >= 30 {
	// 		break
	// 	}
	// 	endIndex++
	// }
	// min2 := costs[2] + mincostTickets(days[endIndex+1:], costs)
	//
	// cost := min(min0, min1, min2)
	// return cost
}

// @leet end

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15}))
}
