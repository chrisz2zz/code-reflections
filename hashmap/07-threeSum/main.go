package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)

	idx := make(map[[3]int]struct{})

	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				idx[[3]int{nums[i], nums[l], nums[r]}] = struct{}{}
			}
			if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	for k := range idx {
		res = append(res, []int{k[0], k[1], k[2]})
	}

	return res
}
