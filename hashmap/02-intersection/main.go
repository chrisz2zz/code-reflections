package main

import (
	"fmt"
)

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersection(nums1 []int, nums2 []int) []int {
	idx := make(map[int]int, len(nums1))

	res := []int{}

	unique := make(map[int]struct{})

	for _, v := range nums1 {
		idx[v]++
	}

	for _, v := range nums2 {
		if idx[v] != 0 {
			unique[v] = struct{}{}
		}
	}

	for v := range unique {
		res = append(res, v)
	}

	return res
}
