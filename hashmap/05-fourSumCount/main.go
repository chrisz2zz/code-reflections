package main

import "fmt"

func main() {
	fmt.Println(fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	idx1 := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			idx1[nums1[i]+nums2[j]]++
		}
	}

	count := 0

	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			if v := idx1[-(nums3[i]+nums4[j])]; v > 0 {
				count += v
			}
		}
	}

	return count
}
