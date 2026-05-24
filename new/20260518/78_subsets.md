# 78. Subsets

## Link

[Leetcode](https://leetcode.com/problems/subsets/description/)

## Difficulty

Medium

## Problem

Given an integer array nums of unique elements, return all possible subsets (the power set) .

The solution set must not contain duplicate subsets. Return the solution in any order .

Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:

Input: nums = [0]
Output: [[],[0]]

Constraints:

1

-10

All the numbers of nums are unique .

## Solve Date

20260524

## Solution

[Code](../../src/78_subsets/main.go)

## Status

DONE in 1 hour, I wake up early today ^^

## Core idea

- recursion at heart, not sure how to explain the intuition correctly, but for each recurse, we need to determine the "decision" we need to make, here its include or skip the current index, and have a base case for when we need to record the result, here is when the index is equal to the array's length
- pay attention to go memory behavior, create a copy when add the sub to the result, or it can be messed

## Failure

- need hints, but it's my first attempt at the pattern, who cares :))

## Success

- quite smooth solve, take me not much effort to understand the intuition

## Tags

`backtracking` `recursion`
