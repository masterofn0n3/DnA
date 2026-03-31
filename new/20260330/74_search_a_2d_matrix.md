# 74. Search a 2D Matrix

## Link

[Leetcode](https://leetcode.com/problems/search-a-2d-matrix/description/)

## Solve Date

20260331

## Solution

[Code](../../src/74_search_a_2d_matrix/main.go)

## Status

DONE

## Core idea

- binary search 2 times, one for the inner array, then fetch the result for the outer to continue binary search
- optimize to deduct real value by virtual index
  - get right pointer right = m \* n
  - mid = (left+right)/2
  - row = mid/numCol
  - col = mid%numCol

```
1 2 3 4
5 6 7 8
9 10 11 12

left = 0
right = 11
mid = 5
row = 5/4 = 1
col = 5%4 = 1

left = 6
right = 11
mid = 8
row = 8/4 = 2
col = 8%4 = 0

```

## Failure

- the loop condition should be left <= right, but I set it to left < right, skip the case of equal

## Success

- No hint needed
- Done by one strike in like 15 mins
- I'm so rock

## Tags

`binary-search`
