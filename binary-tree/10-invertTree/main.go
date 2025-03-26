package main

import "fmt"

func main() {
	p := invertTree(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	})
	display(p)
}

func display(p *TreeNode) {
	if p == nil {
		return
	}

	pp := []*TreeNode{p}

	for len(pp) > 0 {
		length := len(pp)
		for i := range length {
			fmt.Println(pp[i].Val)
			if pp[i].Left != nil {
				pp = append(pp, pp[i].Left)
			}

			if pp[i].Right != nil {
				pp = append(pp, pp[i].Right)
			}
		}
		pp = pp[length:]
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
