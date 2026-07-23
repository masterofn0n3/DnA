# 309. Best Time to Buy and Sell Stock with Cooldown

## Link

[Leetcode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/)

## Difficulty

Medium

## Problem

You are given an array prices where prices[i] is the price of a given stock on the i th day.

Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:

After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Example 1:

Input: prices = [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]

Example 2:

Input: prices = [1]
Output: 0

Constraints:

1

0

## Solve Date

20260722

## Solution

[Code](../../src/309_best_time_to_buy_and_sell_stock_with_cooldown/main.go)

## Status

DONE

## Core idea

- 3 arrays, to store the max value of the profit if we hold, sold, or rest on day i
- to hold[i] that means we can hold[i-1] yesterday and do nothing today, or just rest yesterday and buy today `hold[i] = max(hold[i-1], rest[i-1]-prices[i])`
- to sold[i] that means we have to already hold at yesterday => `sold[i] = hold[i-1] + price[i]`
- to rest[i] that means we already rest yesterday and do nothing today, or must rest because we sold yesterday => `rest[i] = max(rest[i-1], sold[i-1])`

## Failure

wow, which genius could be to solve this at first sight without any hint, definitely not me :)

## Success

nothing when well, beside it got quite intuitive when I understand it, this is called a state machine

## Tags

`2d-dp`
