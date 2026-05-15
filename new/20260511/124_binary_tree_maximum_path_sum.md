# 124. Binary Tree Maximum Path Sum

## Link

[Leetcode](https://leetcode.com/problems/binary-tree-maximum-path-sum/description/)

## Difficulty

Hard

## Problem

A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once . Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path .

Example 1:

Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

Example 2:

Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

Constraints:

The number of nodes in the tree is in the range [1, 3 * 10 4 ] .

-1000

## Solve Date

20260515

## Solution

[Code](../../src/124_binary_tree_maximum_path_sum/main.go)

## Status

DONE quite fast for a hard tho

## Core idea

- dfs
- keep a global max sum, update it if any var is better
- recurse left for left sum, right for right sum
- a path is consider a path if it form a arch, so we couldn't return root + left + right, but need to return root + left or root + right, depend on which is greater
- if left or right is negative we don't add them at all, could eliminate them by clumps them to 0 by `max(return value, 0)`
- initial the maxSum to math.MinInt to avoid any negative skipping

## Failure

- still need hint, but it's a hard anyway

## Success

- i understand the hint quite fast, and also apply the recursive like a instinct haha

## Tags

`tree` `recursive`
