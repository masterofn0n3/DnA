package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(dfs(root.Left), 0)
		right := max(dfs(root.Right), 0)

		sum := root.Val + left + right

		if sum > maxSum {
			maxSum = sum
		}

		return max(root.Val+left, root.Val+right)
	}

	dfs(root)

	return maxSum
}
