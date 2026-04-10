# 239. Sliding Window Maximum

## Link

[Leetcode](https://leetcode.com/problems/sliding-window-maximum/description/)

## Solve Date

20260409

## Solution

[Code](../../src/239_sliding_window_maximum/main.go)

## Status

DONE - kinda intuitive solution for a "hard" problem but I still need hint tho

## Core idea

- thought approach: store the maximum of the first window, when ever the window advance, if the new element is greater than max, replace it, the issue is when the max got pooped of of the window, we need to know the second max, and we need a `deque` for it
- for each iteration
  - we do a for loop to compare the new element with the deque last element, if it's greater than the deque tail, we remove it, by that approach, we effectively remove the elements older and smaller element that couldn't be a max
  - then we append the new element
  - after than compare the queue head value with current i - k, if it's equal, that means the window already pass it, so we need to remove it
  - another if to compare i with k - 1, if i > k -1 means the window had form, and we should start recording by append the queue head to the result

## Failure

- still need tips, and I think I should use a stack to store max, but should use deque, but I have not ever use a deque lmao, not sure why the neetcode roadmap doesn't have a section for it

## Success

- At least I know a new pattern now, still need to learn it tho

## Tags

`slicing window` `queue`
