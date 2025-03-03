package main

import "fmt"

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))

	for i := 0; i < len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i < len(coins); i++ {
		dp[i][0] = 1
	}

	for j := coins[0]; j <= amount; j++ {
		dp[0][j] += dp[0][j-coins[0]]
	}

	for i := 1; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			}
		}
	}

	return dp[len(coins)-1][amount]
}