# 4. Median of Two Sorted Arrays

## Link

[Leetcode](https://leetcode.com/problems/median-of-two-sorted-arrays/description/)

## Solve Date

20260402

## Solution

[Code](../../src/4_median_of_two_sorted_arrays/main.go)

## Status

DONE with so much asking for tips. so hard

## Core idea

- first need to understand what's a median in a sorted array is, in a odd length array, it's the middle element, in a even one, it's the half of the sum of 2 middle elements
- use binary search
  - find the shorter length array, operate on it, if we use the longer array, it could have the chances of we took negative element from the shorter array, adding unnecessary complexity
  - the merge array will have length m + n => so each half is (m+n)/2 => the partition of each array in each half is varies
  - that boundary of partition is what we use binary search for
  - if we take i from array 1 => we need to take j = (m+n+1)/2 from array 2
  - for a valid merge, all the element on the left will always smaller on the elements on the right
    - so the last element of array1 first partition should be smaller than the first element of array2 second partition => if not, we took too much, decrase it by set the right pointer to mid - 1
    - similarly, if last elment of array 2 first half is greater than first element of array 1 last half, we took too much from the array 2, reduce it by increase the left pointer to mid + 1
    - if both above condition are met, calculate the median using max(left half) and min(half right)
    - aware of when we don't took any element from array1, or took all of the element, then the array access later could be runtime panic, use Math.MaxInt and Math.MinInt to safeguard and set the index first before access the array

## Failure

- too hard, feel like i still couldn't wrap my head around this

## Success

- well i guess i conquered some of this problem, yay
- and know a new way to find the median without merge 2 array

## Tags

`binary-search`
