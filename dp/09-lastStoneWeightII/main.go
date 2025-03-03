package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}

func lastStoneWeightII(stones []int) int {
	sum := 0

	for _, v := range stones {
		sum += v
	}

	target := sum / 2

	dp := make([][]int, len(stones))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}

	// 初始化 用第一个石头时,不同背包容量下可以装下的石头重量
	for j := stones[0]; j <= target; j++ {
		dp[0][j] = stones[0]
	}

	for i := 1; i < len(stones); i++ {
		for j := 1; j <= target; j++ {
			if j < stones[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
			}
		}
	}

	return (sum - dp[len(stones)-1][target]) - dp[len(stones)-1][target]  // 大的一堆减小的一堆
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
