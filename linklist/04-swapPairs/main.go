package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		// Next: &ListNode{
		// 	Val: 2,
		// 	Next: &ListNode{
		// 		Val: 3,
		// 		Next: &ListNode{
		// 			Val: 4,
		// 		},
		// 	},
		// },
	}

	h := swapPairs(head)
	for h != nil {
		fmt.Println((h.Val))
		h = h.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	h := pre
	for head != nil {
		tail := head
		for i := 0; i < 2; i++ {
			if tail == nil {
				return pre.Next
			}

			tail = tail.Next
		}

		newHead, newTail := reverseList(head, tail)

		h.Next = newHead
		newTail.Next = tail
		h = newTail
		head = tail
	}

	return pre.Next
}

func reverseList(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := &ListNode{}
	newTail := head
	for head != tail {
		pre.Next, head.Next, head = head, pre.Next, head.Next
	}

	return pre.Next, newTail
}
