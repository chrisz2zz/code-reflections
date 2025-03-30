package main

import "fmt"

func main() {
	fmt.Println(findBottomLeftValue(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	p := []*TreeNode{root}
	pre := []*TreeNode{}

	for len(p) > 0 {
		length := len(p)
		for i := 0; i < length; i++ {
			if p[i].Left != nil {
				p = append(p, p[i].Left)
			}
			if p[i].Right != nil {
				p = append(p, p[i].Right)
			}
		}
		pre = p[:length]
		p = p[length:]
	}

	return pre[0].Val
}
