package main

import "fmt"

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
}

func maxProfit(k int, prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, k*2+1)
	}

	for i := 1; i < k*2+1; i++ {
		if i%2 != 0 {
			dp[0][i] = -prices[0]
		}
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j < 2*k+1; j++ {
			if j%2 != 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}

	return dp[len(prices)-1][2*k]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
