# 200. Number of Islands

## Link

[Leetcode](https://leetcode.com/problems/number-of-islands/description/)

## Difficulty

Medium

## Problem

Given an m x n 2D binary grid grid which represents a map of '1' s (land) and '0' s (water), return the number of islands .

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

 Example 1: 

 Input: grid = [
 ["1","1","1","1","0"],
 ["1","1","0","1","0"],
 ["1","1","0","0","0"],
 ["0","0","0","0","0"]
]
 Output: 1

 Example 2: 

 Input: grid = [
 ["1","1","0","0","0"],
 ["1","1","0","0","0"],
 ["0","0","1","0","0"],
 ["0","0","0","1","1"]
]
 Output: 3

Constraints: 

 m == grid.length 

n == grid[i].length 

1 

grid[i][j] is '0' or '1' .

## Solve Date

20260704

## Solution

[Code](../../src/200_number_of_islands/main.go)

## Status

DONE

## Core idea

- same idea as [word search](), just minus the backtracking

## Failure


## Success

- first problem of graph, huray

## Tags

`graph`
