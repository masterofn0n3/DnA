package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// k = 2
// 1 2 3 4 5 6
func reverseKGroup(head *ListNode, k int) *ListNode {
	var current, groupTail *ListNode
	dummy := &ListNode{}
	groupTail = dummy
	current = head
	for current != nil {
		temp := current
		for i := 0; i < k; i++ {
			if temp == nil {
				return dummy.Next
			}
			temp = temp.Next
		}
		var prev, next *ListNode
		// prev = 2
		// current 3
		newTail := current
		// newTail 1
		for i := 0; i < k; i++ {
			next = current.Next
			current.Next = prev
			prev = current
			current = next
		}
		//after this for loop 1 <= 2
		groupTail.Next = prev
		// dummy.Next = 2
		newTail.Next = current
		groupTail = newTail
	}
	return dummy.Next
}
