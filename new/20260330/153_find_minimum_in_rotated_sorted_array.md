# 153. Find Minimum in Rotated Sorted Array

## Link

[Leetcode](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/)

## Solve Date

20260331

## Solution

[Code](../../src/153_find_minimum_in_rotated_sorted_array/main.go)

## Status

DONE

## Core idea

- binary search left = 0 right = len -1 with condition left < right
- if num[mid] > num[right] that means the break point is on the right side => left = mid + 1
- else the break point is on left side, mid included => right = mid
- left and right will eventually meet
- return left
- O(log n)

## Failure

- i need tips on the solutions
- still vague on the condition and the left and right reassigning

## Success

- kinda easy

## Tags

`binary-search`
