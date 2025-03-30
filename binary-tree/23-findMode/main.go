package main

import "fmt"

func main() {
	fmt.Println(findMode(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	var pre *TreeNode
	max, count := 1, 1
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}

		travel(node.Left)
		if pre != nil && node.Val == pre.Val {
			count++
		} else {
			count = 1
		}

		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}

			max = count
		}
		pre = node

		travel(node.Right)
	}

	travel(root)

	return res
}
