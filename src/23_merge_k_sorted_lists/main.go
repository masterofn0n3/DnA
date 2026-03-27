package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	firstHalf := mergeKLists(lists[:mid])
	secondHalf := mergeKLists(lists[mid:])
	return merge2Lists(firstHalf, secondHalf)
}

func merge2Lists(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy

	for head1 != nil || head2 != nil {
		var node *ListNode
		if head1 == nil {
			prev.Next = head2
			return dummy.Next
		}
		if head2 == nil {
			prev.Next = head1
			return dummy.Next
		}
		if head1.Val <= head2.Val {
			node = head1
			head1 = head1.Next
		} else {
			node = head2
			head2 = head2.Next
		}
		prev.Next = node
		prev = node
	}
	return dummy.Next
}
