# 131. Palindrome Partitioning

## Link

[Leetcode](https://leetcode.com/problems/palindrome-partitioning/description/)

## Difficulty

Medium

## Problem

Given a string s , partition s such that every substring of the partition is a palindrome . Return all possible palindrome partitioning of s .

Example 1:

Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Example 2:

Input: s = "a"
Output: [["a"]]

Constraints:

1

s contains only lowercase English letters.

## Solve Date

20260601

## Solution

[Code](../../src/131_palindrome_partitioning/main.go)

## Status

DONE in a tad difficult

## Core idea

- from index 0, generate a for loop and slice string into index, increase by 1 each interation
- then in the loop, check if the substring is palindrom, if yes, recursive call on it
- finish condition is when index == len(s)

## Failure

- i mistake implement it in [subset](../20260518/78_subsets.md) direction, but its not, got me lost for a bit, and i basically read hint for all of the solution

## Success

## Tags

`backtracking` `recursion`
