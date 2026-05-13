# 105. Construct Binary Tree from Preorder and Inorder Traversal

## Link

[Leetcode](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/)

## Difficulty

Medium

## Problem

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree .

Example 1:

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]

Constraints:

1

inorder.length == preorder.length

-3000

preorder and inorder consist of unique values.

Each value of inorder also appears in preorder .

preorder is guaranteed to be the preorder traversal of the tree.

inorder is guaranteed to be the inorder traversal of the tree.

## Solve Date

20260513

## Solution

[Code](../../src/105_construct_binary_tree_from_preorder_and_inorder_traversal/main.go)

## Status

DONE in 2 distracted days, and a few hours of switching back and forth between solving and scrolling facebook

## Core idea

- 2 list, the preorder to get the root, the inorder to know which side belong to the left and right, use a map to store the value postion in the inorder array, => from there we can deduce the left and right portion, => and recurse it until there's no element in the portion
- a lot of index bookeeping

## Failure

- too many moving part makes it difficult to track, take me a lot of tries and a lot of hints from AI mentor :))

## Success

- i know what inorder, preorder and postorder is and what's there's purpose, in is left root right, pre is root left right, post is left right root

## Tags

`tree` `recursion`
