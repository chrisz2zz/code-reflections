package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	idx := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if k, ok := idx[nums[i]]; ok {
			return []int{k, i}
		}

		idx[target-nums[i]] = i
	}

	return nil
}
