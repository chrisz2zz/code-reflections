package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findTargetSumWays([]int{0}, 0))
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 判断能不能组成
	if int(math.Abs(float64(target))) > sum {
		return 0
	}

	/*  要装满的背包重量
	x-y = target
	x+y = sum
	x-(sum-x) = target
	x = (target + sum) / 2
	*/

	// 左右差值不是个整数
	if (target+sum)%2 != 0 {
		return 0
	}

	bagSize := (target + sum) / 2

	dp := make([][]int, len(nums)) // dp[i][j] 代表0..i个物品装满j容量的背包的方法
	// 特殊情况 nums里的值为0时,随着物品的增加,形成的组合数为2的0的个数次方
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, bagSize+1)
	}

	// 初始化第一行
	if nums[0] <= bagSize {
		dp[0][nums[0]] = 1
	}

	dp[0][0] = 1

	// 初始化最左列
	numZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			numZero++
		}
		dp[i][0] = int(math.Pow(2.0, float64(numZero)))
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= bagSize; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}

	return dp[len(nums)-1][bagSize]
}
