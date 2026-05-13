package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]
/*
	3
9 		15 20 7

*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	inMap := make(map[int]int)
	for i, v := range inorder {
		inMap[v] = i
	}

	var buildTreeRecursive func(prestart, preend, instart, inend int) *TreeNode

	buildTreeRecursive = func(prestart, preend, instart, inend int) *TreeNode {
		if prestart > preend {
			return nil
		}
		currentRoot := &TreeNode{
			Val: preorder[prestart],
		}

		currentRootIndex := inMap[currentRoot.Val]

		leftInStart := instart
		leftInEnd := currentRootIndex - 1
		rightInStart := currentRootIndex + 1
		rightInEnd := inend

		leftPreStart := prestart + 1
		leftPreEnd := leftPreStart + (leftInEnd - leftInStart)
		rightPreStart := leftPreEnd + 1
		rightPreEnd := preend

		currentRoot.Left = buildTreeRecursive(leftPreStart, leftPreEnd, leftInStart, leftInEnd)
		currentRoot.Right = buildTreeRecursive(rightPreStart, rightPreEnd, rightInStart, rightInEnd)

		return currentRoot
	}

	return buildTreeRecursive(0, len(preorder)-1, 0, len(inorder)-1)

}
