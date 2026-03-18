# 84. Largest Rectangle in Histogram

## Link

[Leetcode](https://leetcode.com/problems/largest-rectangle-in-histogram/description/)

## Solve Date

20260312

## Solution

[Code](../../src/84_largest_rectangle_in_histogram/main.go)

## Status

DONE while as was a little bit drunk

## Core idea

- for each i, we can expand the rectangle both the left and right until we met with a column whose height is less than the current column
- use a stack to store the next smaller column for both direction
- then `width = (right-left-1)*height)`
- for element that we couldn't find next lower left and right, need to replace them with len(heights) -1 , and -1 to apply for the formula above
- furthur optimize by add 0 at the end to make every column greater than it, remove the clean up pass
- could do in once pass instead of 4 pass like I do //todo rewrite in one pass in the next review
- O(n)

## Failure

- take times and not in one session because I got distracted with other things
- update other problem .md file instead of the correct one, because still drunk haah

## Success

- manage to think of a solution with just one hint, and I know the correct approach right away, pass beyond just "know that i need to use stack", and this is a hard problem, hehe

## Tags

`array` `stack` `high-priority-for-review`
