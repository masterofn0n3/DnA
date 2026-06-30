# 300. Longest Increasing Subsequence

## Link

[Leetcode](https://leetcode.com/problems/longest-increasing-subsequence/description/)

## Difficulty

Medium

## Problem

Given an integer array nums , return the length of the longest strictly increasing subsequence .

Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:

1

-10 4 4

Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?

## Solve Date

20260630

## Solution

[Code](../../src/300_longest_increasing_subsequence/main.go)

## Status

DONE

## Core idea

- two approach
  - 1
    - dp[i] represent longest subsequent up till nums[i]
    - dp[i] = dp[j] + 1 with j span from 0 to i, with dp[i] is the longest subsequent till j
    - O(n\*2)
  - 2
    - create a tail array, each element i in tail, represent smallest ending value of subsequence of length i
    - tail[0] = nums[0]
    - if num[i] > tail's last element, append to tail
    - else find the index in tail where the element were greater than num[i], and replace it with num[i], using binary search `sort.SearchInt`
    - O(nlogn)

## Failure

- still need hint

## Success

- understand 2 approach

## Tags

`1d-dp`
