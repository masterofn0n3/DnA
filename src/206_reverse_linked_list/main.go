package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var currentNode *ListNode
	var next *ListNode
	for currentNode = head; currentNode != nil; currentNode = next {
		next = currentNode.Next
		currentNode.Next = prev
		prev = currentNode
	}
	return prev
}
