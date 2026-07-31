# 329. Longest Increasing Path in a Matrix

## Link

[Leetcode](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/)

## Difficulty

Hard

## Problem

Given an m x n integers matrix , return the length of the longest increasing path in matrix .

From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).

Example 1:

Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9] .

Example 2:

Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Output: 4
Explanation: The longest increasing path is [3, 4, 5, 6] . Moving diagonally is not allowed.

Example 3:

Input: matrix = [[1]]
Output: 1

Constraints:

m == matrix.length

n == matrix[i].length

1

0 31 - 1

## Solve Date

20260731

## Solution

[Code](../../src/329_longest_increasing_path_in_a_matrix/main.go)

## Status

DONE

## Core idea

- same approach of dfs everycell, return 0 if invalid, dfs + 1 if valid, add memo[[i, j]] to reduce time

## Failure

- some typo, and edge cases

## Success

- too farmiliar with this type of problem

## Tags

`dp`
