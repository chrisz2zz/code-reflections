package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) // 前一天持有股票时现金,前一天卖掉股票并在今天买入股票时的现金
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) // 前一天不持有股票时现金,前一天持有股票并在今天卖掉的现金
	}

	return dp[len(prices)-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
