package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(11, []int{1,2,3,4,5}))
}

func minSubArrayLen(target int, nums []int) int {
	i := 0
	res := math.MaxInt
	sum := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			sub := j - i + 1
			if sub < res {
				res = sub
			}
			sum -= nums[i]
			i++
		}
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}
