package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		subResult := []int{}
		children := []*TreeNode{}
		for len(queue) > 0 {
			current := queue[0]
			subResult = append(subResult, current.Val)
			if current.Left != nil {
				children = append(children, current.Left)
			}
			if current.Right != nil {
				children = append(children, current.Right)
			}
			queue = queue[1:]
		}
		result = append(result, subResult)
		queue = append(queue, children...)
	}
	return result
}

func main()
