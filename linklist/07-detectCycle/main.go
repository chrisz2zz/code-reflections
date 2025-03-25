package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fast = head
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
