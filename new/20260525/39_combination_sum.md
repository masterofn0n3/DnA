# 39. Combination Sum

## Link

[Leetcode](https://leetcode.com/problems/combination-sum/description/)

## Difficulty

Medium

## Problem

Given an array of distinct integers candidates and a target integer target , return a list of all unique combinations of candidates where the chosen numbers sum to target . You may return the combinations in any order .

The same number may be chosen from candidates an unlimited number of times . Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:

Input: candidates = [2], target = 1
Output: []

Constraints:

1

2

All elements of candidates are distinct .

1

## Solve Date

20260526

## Solution

[Code](../../src/39_combination_sum/main.go)

## Status

DONE in 30 mins

## Core idea

- the decision tree here is 1. reuse the same index + include the index in the passing array and 2. pass index + 1(the next index), and skip the current index in the passing array, eventually all the branch will be yield in the call stack
- beware not to include another decision that is `(index+1, [sumArr, candidates[index]])`, which both pass in the next index, and include the current index in the passing array, this branch will eventually appear in decision 2, so if we add this, it will yield duplicated result
- stop and add to result when the passing sum is equal to target, otherwise end the stack when passing sum is greater, or index is out of bound

## Failure

- my solution have a redundant decision branch, detail above

## Success

- i managed to birth out the whole implementation, just missing index out of bound guards and 1 redundant decision branch, but still, I'm awesome dia dia

## Tags

`backtracking` `recursion`
