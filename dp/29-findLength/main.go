package main

import "fmt"

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 4, 5}, []int{3, 2, 1, 4, 7}))
}

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2)+1)
		dp[i][0] = 0
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 0
	}

	res := 0

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
