package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)

	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	fmt.Println(dp)

	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
