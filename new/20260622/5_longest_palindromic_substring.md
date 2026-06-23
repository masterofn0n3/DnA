# 5. Longest Palindromic Substring

## Link

[Leetcode](https://leetcode.com/problems/longest-palindromic-substring/description/)

## Difficulty

Medium

## Problem

Given a string s , return the longest palindromic substring in s .

 Example 1: 

 Input: s = "babad"
 Output: "bab"
 Explanation: "aba" is also a valid answer.

 Example 2: 

 Input: s = "cbbd"
 Output: "bb"

Constraints: 

 1 

s consist of only digits and English letters.

## Solve Date

20260623

## Solution

[Code](../../src/5_longest_palindromic_substring/main.go)

## Status

DONE

## Core idea

- consider each index a valid center, expand from it with 2 case, odd cases then then currrent index is center, even case when current and current + 1 index is center
- write a helper function for the expansion logic
- beware of the cases end > start, when pass in i and i+1 and the value is not equal, so the helper return start < end, so we need a guard for that at the outer func:  end >= start(we have = because `bb` still is a valid cases)

## Failure

- too many index book keeping, hate these type of problem
- this is expand from center approach, I haven't try the dp approach yet

## Success



## Tags

`dp`
