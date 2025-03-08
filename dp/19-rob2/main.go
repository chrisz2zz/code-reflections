package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
}

func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	if len(nums) < 3 {
		return max(nums[0], nums[1])
	}

	// r1 := rob2(nums[1 : len(nums)-1])
	r2 := rob2(nums[:len(nums)-1])
	r3 := rob2(nums[1:])

	// return max(r1, max(r2, r3))
	return max(r2, r3)
}

func rob2(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	if len(nums) < 3 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums)+1)

	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
