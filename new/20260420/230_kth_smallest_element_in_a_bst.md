# 230. Kth Smallest Element in a BST

## Link

[Leetcode](https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/)

## Difficulty

Medium

## Problem

Given the root of a binary search tree, and an integer k , return the k th smallest value ( 1-indexed ) of all the values of the nodes in the tree .

Example 1:

Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

Constraints:

The number of nodes in the tree is n .

1 4

0 4

Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

## Solve Date

20260511

## Solution

[Code](../../src/230_kth_smallest_element_in_a_bst/main.go)

## Status

DONE in one distracted hour, create last month but finish nearly one month later, because I travel accross VN haha

## Core idea

- DFS, because it's a binary tree, traverse in left => root => right in recursion will provide a list of incrementing values
- keep a counter to compare with k, once equal, return that value
- keep a guard at the start to early return to not go through the whole tree

## Failure

- need some small hint for the approach and the optimization

## Success

- i still consider it goood because just need some small hint after few weeks not touching LC, that means I'm retain something in my stupid brain haha

## Tags

`tree` `bfs`
