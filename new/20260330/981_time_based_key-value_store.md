# 981. Time Based Key-Value Store

## Link

[Leetcode](https://leetcode.com/problems/time-based-key-value-store/description/)

## Solve Date

20260401

## Solution

[Code](../../src/981_time_based_key-value_store/main.go)

## Status

DONE

## Core idea

- save value and timestamp for each key into a struct and push them to a slices
- binary search the slices to find the latest timestamp
- pay attention to the left right condition, it should be left <= right, along with reassign left = mid + 1, right = mid - 1, and a value to store the current result when entries[mid] <= timestampt
  - normal right = mid - 1 and left = mid don't work here because in this cases `right` have to advance by mid + 1, not left
  - but if we only set left = mid, and left < right, it will cause infinite loop, because when left + 1 = right, (left + right) /2 = left, left never changes again => inf loop
  - so we need to also advance left by left = mid + 1
  - but if we do that, it will skip the last `mid` when `left` = `right`, because the loop stopped by then, so we need to change the loop to `left <= right`
  - but when we do that, when the loop stop left = right + 1, one pointer ahead of the correct index,
  - remove all the cognitive decision by use a value to store best result and save to it every time entry.timestampt <= timestampt

## Failure

- i still not fully understand the decision needed for the loop condition, so many edge cases

## Success

- but I find the one template to rule all cases, describe more at [binary search](../../pattern/binary_search.md)

## Tags

`binary-search`
