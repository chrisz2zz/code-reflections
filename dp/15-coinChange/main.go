package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}

func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 0
		for j := 1; j <= amount; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	for j := 1; j <= amount; j++ {
		if j%coins[0] == 0 {
			dp[0][j] = j / coins[0]
		}
	}

	for i := 1; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coins[i] && dp[i][j-coins[i]] != math.MaxInt {
				dp[i][j] = min(dp[i][j-coins[i]]+1, dp[i-1][j])
			}
		}
	}

	if dp[len(coins)-1][amount] == math.MaxInt {
		return -1
	}

	return dp[len(coins)-1][amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
