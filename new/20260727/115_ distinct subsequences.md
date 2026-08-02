# 115. Distinct Subsequences

## Link

[Leetcode](https://leetcode.com/problems/distinct-subsequences/description/)

## Difficulty

Hard

## Problem

Given two strings s and t, return the number of distinct subsequences of s which equals t.

The test cases are generated so that the answer fits on a 32-bit signed integer.

 Example 1: 

 Input: s = "rabbbit", t = "rabbit"
 Output: 3
 Explanation: 
As shown below, there are 3 ways you can generate "rabbit" from s.
 rabb b it 
 ra b bbit 
 rab b bit 

 Example 2: 

 Input: s = "babgbag", t = "bag"
 Output: 5
 Explanation: 
As shown below, there are 5 ways you can generate "bag" from s.
 ba b g bag 
 ba bgba g 
 b abgb ag 
 ba b gb ag 
 babg bag 

 Constraints: 

 1 

s and t consist of English letters.

## Solve Date

20260802

## Solution

[Code](../../src/115_ distinct subsequences/main.go)

## Status

DONE

## Core idea

- solved it in 3 levels of optimization
- first one is top down with recursion and memoization, count += dfs(i+1, j) + dfs(i+1, j+1)
- next to convert it to bottom up with recurrence of dp[i][j] = dp[i+i][j]+dp[i+1][j+1]
- because dp[i][j] only depend on row i+1, so we can compressed the 2d array to 1d with array[j] += array[j+1]

## Failure

<!-- What went wrong / what to avoid next time -->

## Success

<!-- What went well / what to reuse -->

## Tags

<!-- e.g. array, hash-set, sorting -->
