# 322. Coin Change

## Link

[Leetcode](https://leetcode.com/problems/coin-change/description/)

## Difficulty

Medium

## Problem

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount . If that amount of money cannot be made up by any combination of the coins, return -1 .

You may assume that you have an infinite number of each kind of coin.

 Example 1: 

 Input: coins = [1,2,5], amount = 11
 Output: 3
 Explanation: 11 = 5 + 5 + 1

 Example 2: 

 Input: coins = [2], amount = 3
 Output: -1

 Example 3: 

 Input: coins = [1], amount = 0
 Output: 0

Constraints: 

 1 

1 31 - 1 

0 4

## Solve Date

20260627

## Solution

[Code](../../src/322_coin_change/main.go)

## Status

DONE 

## Core idea

- take dp where dp[i] is smallest amount of coin needed to reach amount i
  - i = 0 need 0 coin
  - i = 1 need 1 1-coin if have, or -1 if not
  - i = 2 need 1 2-coin if have, or 2 1-coins, or -1 if none
- so, we need to loop through coins array, and set dp[i] = min(dp[i-c] + 1, dp[i]), this works because dp[i-c] also operate on above formula

## Failure

- this approach of use the index number as the amount is new, but cool, I swear till now I still don't fully understand all the use cases of array, the index can be represent in many things

## Success

- new approach loaded

## Tags

`1d-dp`

