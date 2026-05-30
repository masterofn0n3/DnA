# 22. Generate Parentheses

## Link

[Leetcode](https://leetcode.com/problems/generate-parentheses/description/)

## Difficulty

Medium

## Problem

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses .

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

Input: n = 1
Output: ["()"]

Constraints:

1

## Solve Date

20260530

## Solution

[Code](../../src/22_generate_parentheses/main.go)

## Status

DONE

## Core idea

- 2 decision branch, if open > 0, call with `str + "("`, else if close > open call with `str + ")"`, if len of str = with n \* 2, add it to result
- string is immutable, so don't need to backtrack the value

## Failure

- i got the finish condition and branching right, but fail to implemented them, need hint, i modify the str directly in one calls, but i should pass the each new value into its own branching call instead, remember that, no modify in place

## Success

<!-- What went well / what to reuse -->

## Tags

`backtracking`, `recursion`
