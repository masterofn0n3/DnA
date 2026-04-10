# 76. Minimum Window Substring

## Link

[Leetcode](https://leetcode.com/problems/minimum-window-substring/description/)

## Solve Date

20260408

## Solution

[Code](../../src/76_minimum_window_substring/main.go)

## Status

DONE IN 5 HOUR!!!!!

## Core idea

- 2 map, 1 to store `t` char freq, another to for char freq in the window
- keep a count value, increase it when the char freq in window _equal_ to the char require by `t`, decrease it when the char freq in window dip _smaller_ than char required
- keep increase the window while count < number of unique char in `t`
- keep reduce the window while count == number of unique char in `t`, record the best result while reducing

## Failure

- It took me 5 hours since 7 AM to finish, and I still need hint :))), FUCKKKKK

## Success

- at least I'm didn't need to look at any solution
- also I think I learn to pay more attention to the condition, == or > or <= makes a lot difference

## Tags

`sliding-window`
