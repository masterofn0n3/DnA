# 1143. Longest Common Subsequence

## Link

[Leetcode](https://leetcode.com/problems/longest-common-subsequence/description/)

## Difficulty

Medium

## Problem

Given two strings text1 and text2 , return the length of their longest common subsequence . If there is no common subsequence , return 0 .

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde" .

A common subsequence of two strings is a subsequence that is common to both strings.

Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

Constraints:

1

text1 and text2 consist of only lowercase English characters.

## Solve Date

20260721

## Solution

[Code](../../src/1143_longest_common_subsequence/main.go)

## Status

DONE

## Core idea

- dp[i][j] represent LCS of the first i and j char in text1 and text2, so we init dp[len(text1)+1][len(text2)+1]
- because dp[0][j=>n] and dp[i=>n][0] always equal 0, we init the array that way
- nested for loop
- when text1[i] = text2[j], that means we can increase 1 at dp[i][j]
- if not equal, we need to decide which value to set to dp[i][j] to get the max LCS ~ max(dp[i][j-1], dp[i-1][j]) ~ this means decide to skip char from which string
- O(m\*n)

## Failure

- honestly, i don't know if I ever can solve a dp problem by myself

## Success

<!-- What went well / what to reuse -->

## Tags

`2d-dp`
