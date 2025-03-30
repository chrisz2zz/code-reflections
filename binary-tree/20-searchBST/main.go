package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var node *TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	node = nil
	find(root, val)
	return node
}

func find(root *TreeNode, val int) {
	if root == nil {
		return
	}

	if root.Val == val {
		node = root
		return
	}

	find(root.Left, val)
	find(root.Right, val)
}
