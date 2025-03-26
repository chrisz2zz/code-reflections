package main

import "fmt"

func main() {
	fmt.Println(isBalanced(&TreeNode{
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

func isBalanced(root *TreeNode) bool {
	return h(root) != -1
}

func h(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := h(root.Left), h(root.Right)

	if l == -1 || r == -1 {
		return -1
	}

	if l-r > 1 || r-l > 1 {
		return -1
	}

	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
