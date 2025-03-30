package main

import "fmt"

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var root *TreeNode
	max := 0
	maxIdx := 0

	if len(nums) == 0 {
		return root
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			maxIdx = i
			max = nums[i]
		}
	}

	root = &TreeNode{
		Val: max,
	}

	root.Left = constructMaximumBinaryTree(nums[:maxIdx])
	root.Right = constructMaximumBinaryTree(nums[maxIdx+1:])

	return root
}
