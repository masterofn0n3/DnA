# 23. Merge k Sorted Lists

## Link

[Leetcode](https://leetcode.com/problems/merge-k-sorted-lists/description/)

## Solve Date

20260327

## Solution

[Code](../../src/23_merge_k_sorted_lists/main.go)

## Status

DONE

## Core idea

- leverage a merge2List logic
  - create a dummy node, while head1 != nil or head != nil, compare 2 head and append the smaller head to the dummy, and advance the list whose head is appended
  - remember to factor in the case where one of the head is nil, then we append the remaining list to the dummy list
  - O(n) for this
- then use merge sort style recursion to calculate k
  - log k for this
- total time O(n log k)

## Failure

- intuitive solution, but I still need tips to know the solution, and I drag this problem to 2 days to finish, have some emotional issue, but shouldn't let that be an excuse

## Success

<!-- What went well / what to reuse -->

## Tags

`linked-list` `recursion` `merge-sort`
