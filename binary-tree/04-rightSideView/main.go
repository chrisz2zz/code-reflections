package main

import "fmt"

func main() {
	fmt.Println(rightSideView(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	p := []*TreeNode{root}
	for len(p) > 0 {
		length := len(p)
		flag := false
		for i := range length {
			if !flag {
				res = append(res, p[i].Val)
				flag = true
			}

			if p[i].Right != nil {
				p = append(p, p[i].Right)
			}
			if p[i].Left != nil {
				p = append(p, p[i].Left)
			}
		}

		p = p[length:]
	}

	return res
}
