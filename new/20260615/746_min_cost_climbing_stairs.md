# 746. Min Cost Climbing Stairs

## Link

[Leetcode](https://leetcode.com/problems/min-cost-climbing-stairs/description/)

## Difficulty

Easy

## Problem

You are given an integer array cost where cost[i] is the cost of i th step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0 , or the step with index 1 .

Return the minimum cost to reach the top of the floor .

Example 1:

Input: cost = [10, 15 ,20]
Output: 15
Explanation: You will start at index 1.

- Pay 15 and climb two steps to reach the top.
  The total cost is 15.

Example 2:

Input: cost = [ 1 ,100, 1 ,1, 1 ,100, 1 , 1 ,100, 1 ]
Output: 6
Explanation: You will start at index 0.

- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
  The total cost is 6.

Constraints:

2

0

## Solve Date

20260615

## Solution

[Code](../../src/746_min_cost_climbing_stairs/main.go)

## Status

DONE

## Core idea

- break down the problem:
  - create a dp array, dp[i] means the min cost to get to step i
  - each step at i cost a cost[i] amount to _leave_ it
  - if i is 0 or 1, then cost to get there is 0, because we can start right at 0 or 1, no cost
  - if i = 2, means min cost to get to it is `min(dp[1]+cost[1], dp[0]+cost[0])` => deduce to formular `dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])`, we need to compare because, we can choose to step 1 or 2 step at a time

## Failure

- took a while, also I need hint

## Success

- but I manage to figured it out, cheers!

## Tags

<!-- e.g. array, hash-set, sorting -->
