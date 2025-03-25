package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 4)
	}

	dp[0][0] = -prices[0] // 持有股票

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])) // 买入或持有股票
		dp[i][1] = max(dp[i-1][3], dp[i-1][1])                                      // 卖出股票的状态
		dp[i][2] = dp[i-1][0] + prices[i]                                           // 今天卖出股票
		dp[i][3] = dp[i-1][2]                                                       // 冷冻期
	}

	return max(dp[len(prices)-1][2], max(dp[len(prices)-1][3], dp[len(prices)-1][1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
