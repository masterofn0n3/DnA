package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	var dfs func(root *TreeNode, low, high int) bool

	dfs = func(root *TreeNode, low, high int) bool {
		if root == nil {
			return true
		}
		if !(low < root.Val && root.Val < high) {
			return false
		}
		right := dfs(root.Right, root.Val, high)
		left := dfs(root.Left, low, root.Val)

		return right && left
	}

	return dfs(root, math.MinInt, math.MaxInt)

}
