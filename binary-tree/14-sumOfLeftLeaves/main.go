package main

import "fmt"

func main() {
	fmt.Println(sumOfLeftLeaves(&TreeNode{
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

var sum int

func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	findLeft(root)
	return sum
}

func findLeft(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}

	findLeft(root.Left)
	findLeft(root.Right)
}
