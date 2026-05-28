# 46. Permutations

## Link

[Leetcode](https://leetcode.com/problems/permutations/description/)

## Difficulty

Medium

## Problem

Given an array nums of distinct integers, return all the possible permutations . You can return the answer in any order .

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:

Input: nums = [1]
Output: [[1]]

Constraints:

1

-10

All the integers of nums are unique .

## Solve Date

20260528

## Solution

[Code](../../src/46_permutations/main.go)

## Status

DONE in 30 mins

## Core idea

- use a usage array `[len(nums)]bool` to keep track of the usage
- in a recursive function, loop through the usage array, if it's not used, add it to the permutation, change the value to true, recursive call, and then back track

## Failure

- i though i understand the approach, but comes to a new use case, i got no idea at all, still couldn't come up with a full solution

## Success

<!-- What went well / what to reuse -->

## Tags

`backtracking` `recursion`
