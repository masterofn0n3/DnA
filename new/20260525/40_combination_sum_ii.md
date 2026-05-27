# 40. Combination Sum II

## Link

[Leetcode](https://leetcode.com/problems/combination-sum-ii/description/)

## Difficulty

Medium

## Problem

Given a collection of candidate numbers ( candidates ) and a target number ( target ), find all unique combinations in candidates where the candidate numbers sum to target .

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Constraints:

1

1

1

## Solve Date

20260527

## Solution

[Code](../../src/40_combination_sum_ii/main.go)

## Status

DONE in 22min

## Core idea

- same idea with [ComSum I](./39_combination_sum.md)
- the difference is the candidates array has duplicate elements, therefore we had to sort first, and have 2 branch, one is the include branch, the other is the skip branch, before the skip branch, we need to have a loop to increase the index to skip all the duplicated element with current index

## Failure

- I couldn't see through the duplicates issue and how to solve it, need hints

## Success

- I create a approach quite fast(with hint help ofc)

## Tags

`backtracking` `recursion`
