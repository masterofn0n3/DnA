# 130. Surrounded Regions

## Link

[Leetcode](https://leetcode.com/problems/surrounded-regions/description/)

## Difficulty

Medium

## Problem

You are given an m x n matrix board containing letters 'X' and 'O' , capture regions that are surrounded :

Connect : A cell is connected to adjacent cells horizontally or vertically.

Region : To form a region connect every 'O' cell.

Surround : A region is surrounded if none of the 'O' cells in that region are on the edge of the board. Such regions are completely enclosed by 'X' cells.

To capture a surrounded region , replace all 'O' s with 'X' s in-place within the original board. You do not need to return anything.

Example 1:

Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]

Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]

Explanation:

In the above diagram, the bottom region is not captured because it is on the edge of the board and cannot be surrounded.

Example 2:

Input: board = [["X"]]

Output: [["X"]]

Constraints:

m == board.length

n == board[i].length

1

board[i][j] is 'X' or 'O' .

## Solve Date

20260713

## Solution

[Code](../../src/130_surrounded_regions/main.go)

## Status

DONE

## Core idea

- key idea: O cells that at the edge can not be surrounded, and O cells that next to the edged O cells also cannot be surrounded
- dfs at the edge cell, if encounter O => change them to # to marked them, then expand 4 direction, remember guarding
- one last loop to convert the remaining Os to X, and existing # to O
- O(m\*n)

## Failure

- still need a little hint

## Success

- but I managed to solved it by myself, with a small typo (0 instead of O), yay ^^

## Tags

`graph`
