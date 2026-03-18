package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
// output: [1,2,3,5]
/
// Input: head = [1,2,3,4,5], n = 2
// first 1 second 1
// second = 2 => second = 3
// first 2 second 4
// first 3 second 5
// 1 2 - n = 2
// second = nil
//
//
// */
// 1 1
// 1
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first, second := head, head
	for i := 0; i < n && second != nil; i++ {
		second = second.Next
	}
	// meaning after first is the node we have
	if second == nil {
		return first.Next
	}
	for second.Next != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next
	return head

}
