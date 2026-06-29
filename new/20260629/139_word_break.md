# 139. Word Break

## Link

[Leetcode](https://leetcode.com/problems/word-break/description/)

## Difficulty

Medium

## Problem

Given a string s and a dictionary of strings wordDict , return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

 Example 1: 

 Input: s = "leetcode", wordDict = ["leet","code"]
 Output: true
 Explanation: Return true because "leetcode" can be segmented as "leet code".

 Example 2: 

 Input: s = "applepenapple", wordDict = ["apple","pen"]
 Output: true
 Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

 Example 3: 

 Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
 Output: false

Constraints: 

 1 

1 

1 

s and wordDict[i] consist of only lowercase English letters.

All the strings of wordDict are unique .

## Solve Date

20260629

## Solution

[Code](../../src/139_word_break/main.go)

## Status

DONE

## Core idea
- a dp array store if at a index dp[i] whether the substring [0:i] can be segmentation, base case is empty string "" so dp[0] = true, which also required dp have length of len(s)+1
- put dict arr in a map for fast lookup
- check for every i from 0 to len(s)
- then a inner loop to check every j from 0 to i
- if dp[j] = true and s[j:i] exist in dict, which mean s[0:i] can be segmentation, dp[i] = true

## Failure

- why i didn't self solve any of these dp problems

## Success

<!-- What went well / what to reuse -->

## Tags

`1d-dp`
