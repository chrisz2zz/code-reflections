package main

import "fmt"

func main() {
	fmt.Println(levelOrder(&TreeNode{
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

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	p := []*TreeNode{}
	if root != nil {
		p = append(p, root)
	}

	for len(p) != 0 {
		length := len(p)
		tmp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := p[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}

		res = append(res, tmp)
		p = p[length:]
	}

	return res
}
