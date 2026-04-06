# 424. Longest Repeating Character Replacement

## Link

[Leetcode](https://leetcode.com/problems/longest-repeating-character-replacement/description/)

## Solve Date

20260406

## Solution

[Code](../../src/424_longest_repeating_character_replacement/main.go)

## Status

DONE but why so hard

## Core idea

- a map to keep track of the most frequent element, and a value to track the most frequent `maxFreq`
- a value to keep track of best result
- a i, j pointer, form a window, be clear when decide i,j should be inclusive within the window or not, otherwise it will be very confusing(i'm still confusing now)
- traverse the array with j
  - append the frequency to the map, if that frequency greater than `maxFreq`, update it
  - while the window size - maxFrequency > k - reduce the array size by increasing i, and reduce the frequency map of corresponding value
- if (j - i + 1) > best result => update it

## Failure

- i waste 2 hour for this to be not done with the final solution, and have to look up solution in LC for this, this is the first time in this series, so sad(but not sure because of the new modal keep dropping confusing hint or not, because I'm out of limit with Opus 4.6, but I found this new model answer is much longer)
- or just because I'm not focus enough

## Success

<!-- What went well / what to reuse -->

## Tags

`sliding-window`
