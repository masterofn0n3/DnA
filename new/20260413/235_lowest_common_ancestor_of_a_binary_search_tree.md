# 235. Lowest Common Ancestor of a Binary Search Tree

## Link

[Leetcode](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/)

## Solve Date

20260415

## Solution

[Code](../../src/235_lowest_common_ancestor_of_a_binary_search_tree/main.go)

## Status

DONE not that hard for a medium, I guess...

## Core idea

- its a binary tree, so left will always smaller than current node, and right branch will always greater
- so we need to recurse base on these of both p, q smaller/greater than current, then we recurse down that left/right branch, but if p q is separate by two direction, or equal to current node, then current node is LCA

## Failure

- still need hint

## Success

- but I solve quite fast, because the solution is kinda intuitive

## Tags

`tree` `recursion`
