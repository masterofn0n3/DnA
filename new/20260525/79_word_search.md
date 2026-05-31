# 79. Word Search

## Link

[Leetcode](https://leetcode.com/problems/word-search/description/)

## Difficulty

Medium

## Problem

Given an m x n grid of characters board and a string word , return true if word exists in the grid .

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example 1:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Example 2:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Example 3:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

Constraints:

m == board.length

n = board[i].length

1

1

board and word consists of only lowercase and uppercase English letters.

Follow up: Could you use search pruning to make your solution faster with a larger board ?

## Solve Date

20260531

## Solution

[Code](../../src/79_word_search/main.go)

## Status

DONE

## Core idea

- same logic with [word search ii](../20260518/212_word_search_ii.md), but we only have one word here, so simpler

## Failure

- i need hint to short circuit when already found the word, to avoid extra word, also i forget how to mark a cell visited, achieved by set the cell to '#' and backtrack it after done

## Success

- i still remember the approach from word search ii, awesome me

## Tags

`backtracking` `recursion`
