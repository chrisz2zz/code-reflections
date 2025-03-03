package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	dp := make([]int, n+1) // 代表i时所能组成的二叉搜索树数量

	dp[0] = 1 //空树也算一棵树

	// dp[i] = dp[j-1] * dp[i-j]  j-1代表以j为头结点组成的左子树数量,i-j代表以j为头结点组成的右子树数量
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
