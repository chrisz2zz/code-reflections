package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if i >= len(wordDict[j]) {
				if s[i-len(wordDict[j]):i] == wordDict[j] && dp[i-len(wordDict[j])] {
					dp[i] = true
				}
			}
		}
	}

	return dp[len(s)]
}
