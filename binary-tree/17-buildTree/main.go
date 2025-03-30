package main

import "fmt"

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var root *TreeNode

	if len(postorder) == 0 {
		return root
	}

	node := len(postorder) - 1

	root = &TreeNode{
		Val: postorder[node],
	}

	if len(postorder) == 1 {
		return root
	}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[node] {
			break
		}
	}
	// 切割中序数组
	// 左闭右开区间：[0, delimiterIndex)
	leftInorder := inorder[:i]
	// [delimiterIndex + 1, end)
	rightInorder := inorder[i+1:]

	// postorder 舍弃末尾元素
	postorder = postorder[:len(postorder)-1]

	// 切割后序数组
	// 依然左闭右开，注意这里使用了左中序数组大小作为切割点
	// [0, len(leftInorder))
	leftPostorder := postorder[:len(leftInorder)]
	// [len(leftInorder), end)
	rightPostorder := postorder[len(leftInorder):]

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return root
}
