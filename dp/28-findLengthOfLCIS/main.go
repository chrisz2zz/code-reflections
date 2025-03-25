package main

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	res := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}

		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
