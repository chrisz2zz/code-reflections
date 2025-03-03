package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid)) // dp代表0,0到i,j的路径数

	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}

	// 从0,0向下的初始化
	for i := 0; i < len(dp); i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	// 从0,0向右的初始化
	for i := 0; i < len(dp[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	// dp[i][j] = obstacleGrid[i-1][j] 不为1的dp[i-1][j] + obstacleGrid[i][j-1] 不为1的dp[i][j-1]

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
 