# 212. Word Search II

## Link

[Leetcode](https://leetcode.com/problems/word-search-ii/description/)

## Difficulty

Hard

## Problem

Given an m x n board of characters and a list of strings words , return all words on the board .

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

Example 1:

Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]

Example 2:

Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []

Constraints:

m == board.length

n == board[i].length

1

board[i][j] is a lowercase English letter.

1 4

1

words[i] consists of lowercase English letters.

All the strings of words are unique.

## Solve Date

20260522

## Solution

[Code](../../src/212_word_search_ii/main.go)

## Status

DONE in one hour, good problem

## Core idea

- create a trie to store the word
- then create a dfs function(row, col, node) to recursive check every cell, because each word could start from every cell, and each cell could spand in four direction
- beware of edge cases, words overlap, visit duplicated cell in a path, and duplicated words, resolve by continue even when isEnd = true, mark '#' to mark visited cell in a path(remember to reset it when backtrack), and set isEnd = false after add a word to result

## Failure

- need quite a lot of hints

## Success

- but I solve it my self, and I can visualize the solution quite smoothly, for a hard problem

## Tags

`trie` `dfs` `array`
