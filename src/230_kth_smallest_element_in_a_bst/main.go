package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	counter := 0
	value := 0

	var bfs func(root *TreeNode)
	bfs = func(root *TreeNode) {
		if counter >= k {
			return
		}
		if root == nil {
			return
		}
		bfs(root.Left)
		counter++
		if counter == k {
			value = root.Val
			return
		}
		bfs(root.Right)

	}
	bfs(root)

	return value
}
