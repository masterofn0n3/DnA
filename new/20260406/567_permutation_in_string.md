# 567. Permutation in String

## Link

[Leetcode](https://leetcode.com/problems/permutation-in-string/description/)

## Solve Date

20260407

## Solution

[Code](../../src/567_permutation_in_string/main.go)

## Status

DONE

## Core idea

- sliding window, with fixed window size to s1 length, and 2 26 element arrays to store char frequency of 2 string
- advance the window one char per iteration, compare the 2 freq array each time to see if they're match, if yes return true, otherwise modify the freq arr of s2 for the new window position
- 1 loop to form the initital freq arrs, then loop a len(s2)-len(s1) amount to find the match => O(l1 + (l2-l1))

## Failure

- still need hints and need to look at solution

## Success

## Tags

`sliding window`
