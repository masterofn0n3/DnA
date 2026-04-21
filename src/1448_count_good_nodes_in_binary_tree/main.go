package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	counter := 0
	max := math.MinInt

	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		if max <= root.Val {
			max = root.Val
			counter++
		}
		dfs(root.Left, max)
		dfs(root.Right, max)
	}

	dfs(root, max)

	return counter
}
