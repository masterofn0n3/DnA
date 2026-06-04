# 51. N-Queens

## Link

[Leetcode](https://leetcode.com/problems/n-queens/description/)

## Difficulty

Hard

## Problem

The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n , return all distinct solutions to the n-queens puzzle . You may return the answer in any order .

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

Example 1:

Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

Example 2:

Input: n = 1
Output: [["Q"]]

Constraints:

1

## Solve Date

20260604

## Solution

[Code](../../src/51_n-queens/main.go)

## Status

DONE

## Core idea

I'm still not fully understand the idea, but I'll try to explain it

- for queens to not able to attach each other , each row, col, and diagonal can only be occupied by one queen
- we'll take `each one queen each row` as precondition first, and start from there,
- which mean create a queens []int array, each index represent a row, and the value of the index represent a column
- then a recursive function, start from the first row(index), with each we'll check if the cell is valid
  - valid here mean, from the start to the cell, no other queens in the same row(guarantee by the set up), no occupied col(a for loop will do), and no diagonal is occupied too(i still vague at this)
  - if valid, we'll recursive with row + 1
- finish condition is row == n, which mean all row had process, and we got a winner
- then we convert to the result format and add to final result

## Failure

- take me half a day to vaguely understand, need hint ofc

## Success

<!-- What went well / what to reuse -->

## Tags

`need review` `backtracking` `recursion`
