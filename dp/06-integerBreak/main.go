package main

import "fmt"

func main() {
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	dp := make([]int, n+1) // dp[i]代表i所分拆出的数字的最大乘积
	dp[2] = 1              // 第一个有意义的值, 可以分拆,作为初始化

	// dp[i] = max(dp[i], j*(i-j), j*dp[i-j])), dp[i]代表每个j所能构成的最大乘积,j*(i-j)代表两数乘积,j*dp[i-j]代表j乘(i-j所构成的分拆值的最大乘积)
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {  // <=i/2是因为 1*2 2*1会重复计算
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
