# 70. Climbing Stairs

## Link

[Leetcode](https://leetcode.com/problems/climbing-stairs/description/)

## Difficulty

Easy

## Problem

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.

1. 1 step + 1 step
2. 2 steps

Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.

1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:

1

## Solve Date

20260615

## Solution

[Code](../../src/70_climbing_stairs/main.go)

## Status

DONE

## Core idea

- divide the problem into smaller step
- to reach the floor i.e 0, we have 0 way
- to reach step 1, we have 1 way
- to reach step 2, we have 2 ways
- to reach step 3, we have 3 ways = 2 + 1
- to reach step 4, we have 5 ways = 3 + 2
  => deduce that to reach step n , we need step(n-1) + step(n-2)
- this could be achieve by recursive(top down) with memoization to remove the duplicated branch, or bottom up(iterative) to aggregate the total ways from step 0
- iterative solution is more clean and consume less space

## Failure

- in no way this a easy problem, took me 2 distracted days to nearly graps the concept
- at first i approach this in backtracking manner, try every decision, works, but TLE, turn out this little _easy_ problem have a 2 main solution, and 1 optimization

## Success

- i persist and drill it out, so props to me, yayy

## Tags

`1d-dp`
