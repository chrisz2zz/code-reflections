package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	p := []*Node{root}
	for len(p) > 0 {
		length := len(p)

		for i := range length {
			if i > 0 {
				p[i-1].Next = p[i]
			}
			if p[i].Left != nil {
				p = append(p, p[i].Left)
			}

			if p[i].Right != nil {
				p = append(p, p[i].Right)
			}
		}

		p = p[length:]
	}

	return root
}
