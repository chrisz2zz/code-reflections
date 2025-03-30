package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(binaryTreePaths(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res = make([]string, 0)

func binaryTreePaths(root *TreeNode) []string {
	paths(root, "")
	return res
}

func paths(root *TreeNode, s string) {
	if root.Right == nil && root.Left == nil {
		s += strconv.FormatInt(int64(root.Val), 10)
		res = append(res, s)
		return
	}

	s += strconv.FormatInt(int64(root.Val), 10) + "->"
	if root.Left != nil {
		paths(root.Left, s)
	}
	if root.Right != nil {
		paths(root.Right, s)
	}
}
