package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
		return root
	}

	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
		} else {
			insertIntoBST(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: val,
			}
		} else {
			insertIntoBST(root.Left, val)
		}
	}

	return root
}
