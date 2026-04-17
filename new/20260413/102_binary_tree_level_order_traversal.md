# 102. Binary Tree Level Order Traversal

## Link

[Leetcode](https://leetcode.com/problems/binary-tree-level-order-traversal/description/)

## Solve Date

20260416

## Solution

[Code](../../src/102_binary_tree_level_order_traversal/main.go)

## Status

DONE in 30 mins at 7.54 AM because I wake up early haha, now I'm gonna go run, so proud of myself

## Core idea

- Use BFS to traverse the tree
  - a FIFO queue, to store node of current levels, loop until the queue len is 0
  - a inner for loop, also run until the queue len is 0, to append the value of the same level to the result, and append the children of those current level's node
  - after the inner loop is finished, append the children to the queue

## Failure

- need hint, but I never do BFS before so no harsh feelings

## Success

- I know how to implement BFS

## Tags

`tree` `bfs`
