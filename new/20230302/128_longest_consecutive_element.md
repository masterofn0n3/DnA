# 128 Longgest Consecutive Sequence

## Link

[Leetcode](https://leetcode.com/problems/longest-consecutive-sequence/description/)

## Solution

[Code](../../src/128_longest_consecutive_element/main.go)

## Status

DONE

## Core idea

Loop once to transfer the values into a map for quick check existen
Only consider a value v is starting point if v-1 not exist
Then do a inner loop with v+1, v+2,... to get the longest sequence
Then inner loop gothrough each element at most once, so it's not O(n\*2), but O(n)

## Failure

- I need tip to think of a solution, but manage to execute it dependently

## Tags

`array`, `map`
