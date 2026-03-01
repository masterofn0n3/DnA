# 238. Product of Array Except Self

## Link

[Leetcode](https://leetcode.com/problems/product-of-array-except-self/description/)

## Solution

[Code](../../src/238_product_of_array_except_self/main.go)

## Status

DONE in nearly an hour and a half :), but I got distracted too much, thinking hurts my brain

## Core idea

- `output[i]` = (product of everything left of i) × (product of everything right of i).
- Two passes: build prefix product left-to-right, suffix product right-to-left; then combine. O(n) time, O(n) space (or O(1) extra if reusing output array).

## Failure

- think of how to calculate the accumulate value for each index
- Because we can't use a nested loop, think of how to leverage other data structures to store values — in this case, 2 arrays.
- Know the prefix sum/product pattern: use a second array to calculate the accumulated sum (or product) of previous elements. This problem uses the same idea but needs prefix product from both directions., this issue use the same idea, but because we need accumulate product except the current element, prefix product from 2 direction is needed

## Tags

`array`, `prefix sum`
