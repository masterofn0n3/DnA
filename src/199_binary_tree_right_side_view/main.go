package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func rightSideView(root *TreeNode) []int {
// 	result := []int{}
// 	if root == nil {
// 		return result
// 	}

// 	deque := []*TreeNode{root}

// 	for len(deque) > 0 {
// 		result = append(result, deque[len(deque)-1].Val)
// 		children := []*TreeNode{}
// 		for len(deque) > 0 {
// 			current := deque[0]
// 			if current.Left != nil {
// 				children = append(children, current.Left)
// 			}
// 			if current.Right != nil {
// 				children = append(children, current.Right)
// 			}
// 			deque = deque[1:]
// 		}
// 		deque = append(deque, children...)
// 	}

// 	return result
// }

// func rightSideView(root *TreeNode) []int {
// 	result := []int{}
// 	if root == nil {
// 		return result
// 	}

// 	deque := []*TreeNode{root}

// 	for len(deque) > 0 {
// 		levelSize := len(deque)
// 		for i := 0; i <= levelSize-1; i++ {
// 			current := deque[i]
// 			left := current.Left
// 			if left != nil {
// 				deque = append(deque, left)
// 			}
// 			right := current.Right
// 			if right != nil {
// 				deque = append(deque, right)
// 			}
// 			if i == levelSize-1 {
// 				result = append(result, current.Val)
// 			}
// 		}
// 		deque = deque[levelSize:]
// 	}

//		return result
//	}
func rightSideView(root *TreeNode) []int {
	result := []int{}

	var rfs func(node *TreeNode, depth int)
	rfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(result) {
			result = append(result, root.Val)
		}

		rfs(root.Right, depth+1)
		rfs(root.Left, depth+1)
	}

	rfs(root, 0)

	return result
}
