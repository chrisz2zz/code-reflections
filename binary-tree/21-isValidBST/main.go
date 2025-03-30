package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return find(root.Left, root.Val, math.MinInt) && find(root.Right, math.MaxInt, root.Val)
}

func find(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return find(root.Left, root.Val, min) && find(root.Right, max, root.Val)
}
