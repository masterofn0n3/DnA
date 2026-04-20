# 199. Binary Tree Right Side View

## Link

[Leetcode](https://leetcode.com/problems/binary-tree-right-side-view/description/)

## Solve Date

20260420

## Solution

[Code](../../src/199_binary_tree_right_side_view/main.go)

## Status

DONE by a lot of unfocus and procastination

## Core idea

- BFS
  - use a deque to enqueue the children and pop the parent, record the last element in the children array
  - or define a levelSize = len(deque), for loop i = 0 until i upto levelSize-1, inside loop equeue the element, record the last element(when i = levelSize-1), that the right side node, then deque levelSize node from the queue
- DFS
  - create a dfs function, recursion with visiting node.Right first, pass in the depth+1, if depth == len(result), that means that depth is never visited before, if node == nil return, much shorter and cleaner than bfs

## Failure

- I procastinated about 3 days before sit at 1AM to figured the bfs approach by my self, I read hint for DFS first but still couldn't do it, have to read a few more times

## Success

- I solve by bfs DIY, awesome!!!

## Tags

`tree` `bfs` `dfs` `recursion`
