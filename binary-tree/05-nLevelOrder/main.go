package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	p := []*Node{root}

	for len(p) > 0 {
		length := len(p)
		tmp := make([]int, 0)
		for i := range length {
			tmp = append(tmp, p[i].Val)
			for _, v := range p[i].Children {
				if v != nil {
					p = append(p, v)
				}
			}
		}
		res = append(res, tmp)
		p = p[length:]
	}

	return res
}
