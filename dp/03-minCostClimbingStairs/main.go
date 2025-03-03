package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = dp[i-1] + cost[i-1]
		if dp[i] > dp[i-2]+cost[i-2] {
			dp[i] = dp[i-2] + cost[i-2]
		}
	}

	return dp[len(cost)]
}
