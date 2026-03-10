# 739. Daily Temperatures

## Solve date

20260310

## Link

[Leetcode](https://leetcode.com/problems/daily-temperatures/description/)

## Solution

[Code](../../src/739_daily_temperatures/main.go)

## Status

DONE IN MORE THAN AN HOUR

## Core idea

- bruce force by 2 loops that check forward with each index O(n\*2)
- optimize solution using stack that stored the index that hasn't have warmer temperature day yet, then an inner loop that pop and push accordingly O(n)

## Failure

- unfocus, I write code while my head was still in the cloud, causing vague when implementing
- couldn't figure out how to use stack for the optimize solution even when I already know I need to use stack
- nexttime, instead of think "what's next", think: "what was before this"

## Success

<!-- What went well / what to reuse -->

## Tags

`stack` `array` `high-prioritize-for-review`
