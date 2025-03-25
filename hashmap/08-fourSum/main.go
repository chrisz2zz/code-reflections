package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 4, 0, len(nums)-1, target)
}

func nSum(nums []int, num, l, r int, target int) [][]int {
	if num == 2 {
		return twoSum(nums, l, r, target)
	} else {
		res := [][]int{}
		for i := l; i < len(nums); i++ {
			sub := nSum(nums, num-1, i+1, r, target-nums[i])
			for _, v := range sub {
				res = append(res, append([]int{nums[i]}, v...))
			}

			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
		return res
	}
}

func twoSum(nums []int, l, r int, target int) [][]int {
	res := [][]int{}
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			res = append(res, []int{nums[l], nums[r]})
		}

		rr, ll := nums[r], nums[l]
		if sum > target {
			for l < r && nums[r] == rr {
				r--
			}
		} else {
			for l < r && nums[l] == ll {
				l++
			}
		}
	}

	return res
}
