package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)

	queue := []*Node{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if _, ok := nodeMap[current]; !ok {
			nodeMap[current] = &Node{
				Val: current.Val,
			}
		}
		newNei := []*Node{}
		for _, nei := range current.Neighbors {
			if _, ok := nodeMap[nei]; !ok {
				nodeMap[nei] = &Node{
					Val: nei.Val,
				}
				queue = append(queue, nei)
			}
			newNei = append(newNei, nodeMap[nei])
		}
		nodeMap[current].Neighbors = newNei

	}

	return nodeMap[node]
}
