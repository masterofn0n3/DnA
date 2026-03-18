package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

how do I deepclone, by creating a node from scratch
but each node contain 2 pointer
so to deepclone a node, I need to know its next and it's random pointer
do I?

i am given the head
new Node = {val, next, random}
*/
func copyRandomList(head *Node) *Node {

	// 1 2 3
	// 1
	//
	// dummyNode := &Node{}
	// prevNode := dummyNode

	// THIS IS FOR DEEPCLONE A LIST
	// for current := head; current != nil; current = current.Next {
	// 	node := &Node{
	// 		Val: current.Val,
	// 	}
	// 	prevNode.Next = node
	// 	prevNode = node
	// }
	// return dummyNode.Next

	nodeMap := make(map[*Node]*Node)
	for current := head; current != nil; current = current.Next {
		nodeMap[current] = &Node{Val: current.Val}
	}
	for current := head; current != nil; current = current.Next {
		nodeMap[current].Next = nodeMap[current.Next]
		nodeMap[current].Random = nodeMap[current.Random]
	}

	return nodeMap[head]
}
