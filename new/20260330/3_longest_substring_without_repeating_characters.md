# 3. Longest Substring Without Repeating Characters

## Link

[Leetcode](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)

## Solve Date

20260403

## Solution

[Code](../../src/3_longest_substring_without_repeating_characters/main.go)

## Status

DONE

## Core idea

- use a map to store occurence
- and 2 pointer start at the same position, represent a "window"
- advance the right until end of the array
- if the key is not in the map, that means no duplication, increase the right edge, update the length by compare
- if the key exist, delete the key and increase the left edge
- return the length
- O(n)

## Failure

- still need hint on how to increase or decrease the window

## Success

- i still vaguely remember the solution, since I've done this 2 years ago, I guess there's still something stick to me haha

## Tags

`sliding-window`
