package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current1, current2 := l1, l2
	carry := 0
	dummyNode := &ListNode{}
	prev := dummyNode
	for current1 != nil || current2 != nil || carry != 0 {
		var val1, val2 int
		if current1 != nil {
			val1 = current1.Val
		}
		if current2 != nil {
			val2 = current2.Val
		}
		sum := val1 + val2 + carry
		carry = sum / 10
		sum = sum % 10
		prev.Next = &ListNode{Val: sum}
		prev = prev.Next
		if current1 != nil {
			current1 = current1.Next
		}
		if current2 != nil {
			current2 = current2.Next
		}

	}
	return dummyNode.Next
}
