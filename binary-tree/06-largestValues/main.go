package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	p := []*TreeNode{root}

	for len(p) > 0 {
		length := len(p)
		max := math.MinInt
		for i := range length {
			if p[i].Val > max {
				max = p[i].Val
			}

			if p[i].Left != nil {
				p = append(p, p[i].Left)
			}
			if p[i].Right != nil {
				p = append(p, p[i].Right)
			}
		}

		res = append(res, max)
		p = p[length:]
	}

	return res
}
