package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{6, 0, 1, 4, 9, 7, -3, 1, -4, -8, 4, -7, -3, 3, 2, -3, 9, 5, -4, 0}, 6))
}

func topKFrequent(nums []int, k int) []int {
	idx := make(map[int]int)
	for _, item := range nums {
		idx[item]++
	}

	res := make([]int, 0, len(idx))

	for v := range idx {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool {
		return idx[res[i]] > idx[res[j]]
	})

	return res[:k]
}

func topKFrequent2(nums []int, k int) []int {
	ans := []int{}
	map_num := map[int]int{}
	for _, item := range nums {
		map_num[item]++
	}
	for key, _ := range map_num {
		ans = append(ans, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(ans, func(a, b int) bool {
		return map_num[ans[a]] > map_num[ans[b]]
	})
	return ans[:k]
}
