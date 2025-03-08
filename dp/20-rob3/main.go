package main

import (
	"fmt"
	"slices"
)

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	fmt.Println(rob(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	dp := robb(root)
	return slices.Max(dp)
}

func robb(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	left := robb(root.Left)
	right := robb(root.Right)

	robCur := root.Val + left[0] + right[0]

	notRobCur := slices.Max(left) + slices.Max(right)

	return []int{notRobCur, robCur}
}
