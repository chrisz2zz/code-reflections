package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m) // 从0,0到i,j的路径数
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 从0,0往下走的初始化
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// 从0,0往右走的初始化
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	// dp[i][j]=从左边过来+从上面过来
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
