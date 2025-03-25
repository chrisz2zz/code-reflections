package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	for head != nil {
		pre.Next, head.Next, head = head, pre.Next, head.Next
	}

	return pre.Next
}
