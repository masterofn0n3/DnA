# 36. Valid Sudoku

## Link

[Leetcode](https://leetcode.com/problems/valid-sudoku/description/)

## Solution

[Code](../../src/238_product_of_array_except_self/main.go)

## Status

DONE in nearly an hour

## Core idea

- Keep 3 nested array for row, column, boxOccurence
- Use grid index division to find the box index = (i/3)\*3+(j/3)

O(n\*2)

## Failure

- I couldn't think of a solution for this one, so frustuated
- At first I try to valid the whole row and column or at once, by using the current index of the board, which create unnecessary complexity
- Mean while I could decrease the logic by adding more resource to store the occurence(3 2-D array to be extract)
- Need to put this back to review asap

## Tags

`array`, `matrix`
