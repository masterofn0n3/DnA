package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := ""
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			result += "# "
			return
		}
		result += strconv.Itoa(root.Val) + " "
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	tokens := strings.Split(strings.TrimSpace(data), " ")
	if len(tokens) == 0 {
		return nil
	}
	index := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if tokens[index] == "#" {
			index++
			return nil
		}
		val, _ := strconv.Atoi(tokens[index])
		current := &TreeNode{
			Val: val,
		}
		index++
		current.Left = dfs()
		// index++
		current.Right = dfs()

		return current
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
