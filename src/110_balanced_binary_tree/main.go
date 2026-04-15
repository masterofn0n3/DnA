package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := dfs(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := dfs(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}
