package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if subRoot == nil {
		return false
	}
	same := isSame(root, subRoot)
	if same {
		return true
	}

	sameLeft := isSubtree(root.Left, subRoot)
	sameRight := isSubtree(root.Right, subRoot)

	return sameLeft || sameRight

}

func isSame(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root != nil && subRoot == nil || root == nil && subRoot != nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}
	equalLeft := isSame(root.Left, subRoot.Left)
	equalRight := isSame(root.Right, subRoot.Right)

	return equalLeft && equalRight
}
