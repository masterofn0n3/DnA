# 994. Rotting Oranges

## Link

[Leetcode](https://leetcode.com/problems/rotting-oranges/description/)

## Difficulty

Medium

## Problem

You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,

1 representing a fresh orange, or

2 representing a rotten orange.

Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange . If this is impossible, return -1 .

Example 1:

Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4

Example 2:

Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

Example 3:

Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.

Constraints:

m == grid.length

n == grid[i].length

1

grid[i][j] is 0 , 1 , or 2 .

## Solve Date

20260709

## Solution

[Code](../../src/994_rotting_oranges/main.go)

## Status

DONE

## Core idea

- new type of traversal, multi-source traversal
- create a minutes to store time, set to -1, because the first loop doesn't count in time passing
- similar to the islands problem, but this time we don't use dfs
- create a queue to hold all initial rotten cell
- push them all into a queue
- when queue length > 0, create a inner loop to process all CURRENT item in the queue all at once
  - assign a current value to it, pop the queue
  - loop through all four direction, and if encounter fresh cell, rotten it and add to queue, remember to do all the guard checks of invalid position
- after that traverse the grid again, this time find fresh cell, if exist, return -1
- if not return max(0, minutes) // to handle the cases the loop doesn't run

## Failure

<!-- What went wrong / what to avoid next time -->

## Success

- new stuff to learn, yayyy, i'm so sick these last few days, so the progress had been slow, but just keep pusing

## Tags

`graph`
