package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var node *TreeNode
	if root1 == nil && root2 == nil {
		return node
	}

	node = &TreeNode{}
	var r1l, r1r *TreeNode
	if root1 != nil {
		node.Val += root1.Val
		r1l = root1.Left
		r1r = root1.Right
	}

	var r2l, r2r *TreeNode
	if root2 != nil {
		node.Val += root2.Val
		r2l = root2.Left
		r2r = root2.Right
	}

	node.Left = mergeTrees(r1l, r2l)
	node.Right = mergeTrees(r1r, r2r)

	return node
}
