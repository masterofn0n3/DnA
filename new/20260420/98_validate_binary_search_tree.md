# 98. Validate Binary Search Tree

## Link

[Leetcode](https://leetcode.com/problems/validate-binary-search-tree/description/)

## Solve Date

20260422

## Solution

[Code](../../src/98_validate_binary_search_tree/main.go)

## Status

DONE and got fooled again

## Core idea

- dfs, send a interval to each recursive call (low, high), send (current.Val, high) for right child, (low, current.Val) for left child

## Failure

- I though it was easy, but got fooled by the description again
- SHOULD read desc more carefully

## Success

- Figured the approach quite fast

## Tags

- `tree` `dfs`
