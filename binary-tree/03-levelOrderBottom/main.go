package main

import "fmt"

func main() {
	fmt.Println(levelOrderBottom(&TreeNode{
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

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	p := []*TreeNode{}
	if root != nil {
		p = []*TreeNode{root}
	}

	for len(p) > 0 {
		length := len(p)
		tmp := make([]int, 0)
		for i := range length {
			tmp = append(tmp, p[i].Val)
			if p[i].Left != nil {
				p = append(p, p[i].Left)
			}
			if p[i].Right != nil {
				p = append(p, p[i].Right)
			}
		}
		res = append([][]int{tmp}, res...)
		p = p[length:]
	}

	return res
}
