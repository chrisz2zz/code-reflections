package main

import "fmt"

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)


	dp := make([]int, n+1)

	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}

	fmt.Println(dp[n])
	 
}
