# 494. Target Sum

## Link

[Leetcode](https://leetcode.com/problems/target-sum/description/)

## Difficulty

Medium

## Problem

You are given an integer array nums and an integer target .

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

For example, if nums = [2, 1] , you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1" .

Return the number of different expressions that you can build, which evaluates to target .

Example 1:

Input: nums = [1,1,1,1,1], target = 3
Output: 5
Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

Example 2:

Input: nums = [1], target = 1
Output: 1

Constraints:

1

0

0

-1000

## Solve Date

20260727

## Solution

[Code](../../src/494_target_sum/main.go)

## Status

DONE yesterday

## Core idea

- top down with memoization: recursion, and a map to store result to avoid recompute
- i didn't solve using bottom up yet, because i'm a stupid sandwitch

## Failure

- lazy ass bitch me, watch too much tiktok

## Success

- top down is much more intuitive

## Tags

`dp`
