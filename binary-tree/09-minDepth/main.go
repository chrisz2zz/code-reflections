package main

import "fmt"

func main() {
	fmt.Println(minDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return min(minDepth(root.Left)+1, minDepth(root.Right)+1)
}

func min(a, b int) int {
	if a == 1 {
		return b
	}

	if b == 1 {
		return a
	}

	if a < b {
		return a
	}

	return b
}
