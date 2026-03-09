# 150. Evaluate Reverse Polish Notation

## Link

[Leetcode](https://leetcode.com/problems/evaluate-reverse-polish-notation/description/)

## Solution

[Code](../../src/150_evaluate_reverse_polish_notation/main.go)

## Status

DONE

## Core idea

- loop through and append the element to a stack, if encounter a operator then took the first and 2nd element and use them with the found operator, then append the result back to the stack
- O(n)

## Failure

- not focus to solve this simple problem
- been neglect these pass days

## Success

## Tags

`stack` `array`
