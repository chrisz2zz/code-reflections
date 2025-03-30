package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	nums := inorder(root)

	min := math.MaxInt
	for i := 1; i < len(nums); i++ {
		if math.Abs(float64(nums[i]-nums[i-1])) < float64(min) {
			min = nums[i] - nums[i-1]
		}
	}

	return min
}

func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, inorder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorder(root.Right)...)

	return res
}
