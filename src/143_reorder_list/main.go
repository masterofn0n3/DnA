package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3 4 5
// 1 2 3
// 2 4 nil
// 1 2
// 2 4
// 1 2 3 4
// 1 4 2 3
// if node num is even then slow is the not the last node
// if node num is odd then slow is the last node
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// to reverse a list, we need to set
	// current node next to previous node
	// next node next to current
	// 1 2 3 4
	// 1 2 1.next = nil prev = 1
	// 2.next = prev
	//
	var prev, next *ListNode
	currentNode := slow.Next
	slow.Next = nil

	for currentNode != nil {
		next = currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next
	}
	// prev is the first node
	// then we need to set the slow.Next to nil to seperate 2 linkedlist
	// 1 2 3
	// 4 5
	// i need to connect the first head to the second head
	// and connect back to
	// nah I can use a entire new list and connect both from each list for each iteration
	head1, head2 := head, prev
	var next1, next2 *ListNode
	for head1 != nil && head2 != nil {
		next1 = head1.Next
		next2 = head2.Next
		head1.Next = head2
		head2.Next = next1
		head1 = next1
		head2 = next2
	}

}
