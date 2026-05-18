# 297. Serialize and Deserialize Binary Tree

## Link

[Leetcode](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/)

## Difficulty

Hard

## Problem

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree . You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

Example 1:

Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]

Example 2:

Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 10 4 ] .

-1000

## Solve Date

20260518

## Solution

[Code](../../src/297_serialize_and_deserialize_binary_tree/main.go)

## Status

DONE quite smooth for a hard

## Core idea

- for serialize, use preorder traversal root left right to append to the result string, if node is nil then append #
- for serialize, split the input string to tokens, then set a global index = 0, recurse with a helper function to each initiate a node, then set left and right child, each time increase index by 1

## Failure

- i still need hint, but feels like I'm productive today
- still confuse why

```go
if node == nil {
    return nil
}
index++
dfs()//left
index++
dfs()//right
```

and

```go
if node == nil {
    index++
    return nil
}
index++
dfs()//left
dfs()//right
```

work the same

(lie on the bed think about it for few minutes, I guess one consume index on current token, other point to next token)

## Success

- quite intuitive solve, since these pattern appear on previous issue

## Tags

`tree` `recursion`
