# 90. Subsets II

## Link

[Leetcode](https://leetcode.com/problems/subsets-ii/description/)

## Difficulty

Medium

## Problem

Given an integer array nums that may contain duplicates, return all possible subsets (the power set) .

The solution set must not contain duplicate subsets. Return the solution in any order .

Example 1:

Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

Example 2:

Input: nums = [0]
Output: [[],[0]]

Constraints:

1

-10

## Solve Date

20260529

## Solution

[Code](../../src/90_subsets_ii/main.go)

## Status

DONE in 20 mins

## Core idea

- same idea with [subset](../20260518/78_subsets.md), but add the remove duplicating logic of [com sum](40_combination_sum_ii.md), add a guard to check index out of bound in the for loop

## Failure

- i mistakenly put `if` instead of `for` for the loop

## Success

- _almost_ got this without hint, but i got some cases wrong

## Tags

`backtracking` `recursion`
