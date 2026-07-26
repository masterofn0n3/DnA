# 518. Coin Change II

## Link

[Leetcode](https://leetcode.com/problems/coin-change-ii/description/)

## Difficulty

Medium

## Problem

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the number of combinations that make up that amount . If that amount of money cannot be made up by any combination of the coins, return 0 .

You may assume that you have an infinite number of each kind of coin.

The final answer is guaranteed to fit into a signed 32-bit integer.

Example 1:

Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.

Example 3:

Input: amount = 10, coins = [10]
Output: 1

Constraints:

1

1

All the values of coins are unique .

0

## Solve Date

20260726

## Solution

[Code](../../src/518_coin_change_ii/main.go)

## Status

DONE

## Core idea

- dp[i] represent number of combination ways to get to the value `i`
- dp[0] = 1 => because there 1 way to get to value 0, which is use 0 one time
- 2 loop, the outer loop go through each coin denomination, to update the dp array in the inner loop
- number of combination at i equals to itseft + number of combination at [i-c], include itself to preserve the ways of previous coin denomination
- O(m\*n)

## Failure

- my mind is mush right now, spend a lot of time finding new apartment this week, and scroll a bunch of tiktok and facebook while doing that

## Success

## Tags

`dp`
