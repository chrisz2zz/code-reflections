package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}

func fib(n int) int {
	if n < 1 {
		return 0
	}

	dp := make([]int, 2)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[1], dp[0] = dp[0]+dp[1], dp[1]
	}

	return dp[1]
}
