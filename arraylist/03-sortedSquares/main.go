package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	l, r := 0, len(nums)-1

	idx := len(nums) - 1
	for l <= r {
		a := nums[l] * nums[l]
		b := nums[r] * nums[r]
		if a < b {
			res[idx] = b
			r--
			idx--
		} else {
			res[idx] = a
			l++
			idx--
		}
	}

	return res
}
