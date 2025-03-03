package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}

// canPartition 函数检查给定的整数数组是否可以分割成两个和相等的子集
func canPartition(nums []int) bool {
    // 计算数组所有元素的总和
    sum := 0
    for _, v := range nums {
        sum += v
    }

    // 如果总和是奇数，无法分割成两个和相等的子集
    if sum%2 != 0 {
        return false
    }

    // 目标和为总和的一半
    sum /= 2

    // 创建二维DP数组，dp[i][j]表示考虑前i+1个数字能否组成和为j的子集
    // 行数为数组长度，列数为目标和+1
    dp := make([][]int, len(nums))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, sum+1)
    }

    // 初始化第一行：只考虑第一个数字时，能够得到的子集和只有nums[0]
    // 对于所有大于等于nums[0]的目标值j，dp[0][j]都等于nums[0]
    for i := nums[0]; i <= sum; i++ {
        dp[0][i] = nums[0]
    }

    // 填充DP表格
    // 从第二个数字开始考虑（索引为1），因为第一个数字已在上面初始化
    for i := 1; i < len(nums); i++ {
        // j从0开始是因为我们需要考虑所有可能的子集和（0到sum）
        for j := 0; j <= sum; j++ {
            if j < nums[i] {
                // 如果当前目标和j小于当前数字nums[i]，无法选择该数字
                // 直接继承前i-1个数字能否组成和为j的结果
                dp[i][j] = dp[i-1][j]
            } else {
                // 当前目标和j大于等于当前数字nums[i]，有两种选择：
                // 1. 不选择当前数字，继承dp[i-1][j]
                // 2. 选择当前数字，并查看dp[i-1][j-nums[i]]
                // 取这两种情况下能获得的最大和
                dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
            }
        }
    }

    // 检查是否能够得到和为目标值(sum)的子集
    // 如果dp[len(nums)-1][sum]==sum，则表示可以分割成两个和相等的子集
    return dp[len(nums)-1][sum] == sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
