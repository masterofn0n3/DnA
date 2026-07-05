# 695. Max Area of Island

## Link

[Leetcode](https://leetcode.com/problems/max-area-of-island/description/)

## Difficulty

Medium

## Problem

You are given an m x n binary matrix grid . An island is a group of 1 's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid . If there is no island, return 0 .

 Example 1: 

 Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 Output: 6
 Explanation: The answer is not 11, because the island must be connected 4-directionally.

 Example 2: 

 Input: grid = [[0,0,0,0,0,0,0,0]]
 Output: 0

Constraints: 

 m == grid.length 

n == grid[i].length 

1 

grid[i][j] is either 0 or 1 .

## Solve Date

20260705

## Solution

[Code](../../src/695_max_area_of_island/main.go)

## Status

DONE 

## Core idea

- same as 200, dfs and expand with each cell, return a count number 0 or 1 + dfs...

## Failure

- too many typo, and wrong logic to record count

## Success

<!-- What went well / what to reuse -->

## Tags

`graph`
