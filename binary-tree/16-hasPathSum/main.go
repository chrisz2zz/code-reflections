package main

import "fmt"

func main() {
	fmt.Println(hasPathSum(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}, 100))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res bool

func hasPathSum(root *TreeNode, targetSum int) bool {
	res = false
	if root == nil {
		return res
	}
	find(root, targetSum)

	return res
}

func find(root *TreeNode, target int) {
	if root == nil {
		return
	}

	if target-root.Val == 0 && root.Left == nil && root.Right == nil {
		res = true
		return
	}

	if target-root.Val < 0 {
		return
	}

	find(root.Left, target-root.Val)
	find(root.Right, target-root.Val)
}
