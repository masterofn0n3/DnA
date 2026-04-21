# 1448. Count Good Nodes in Binary Tree

## Link

[Leetcode](https://leetcode.com/problems/count-good-nodes-in-binary-tree/description/)

## Solve Date

20260421

## Solution

[Code](../../src/1448_count_good_nodes_in_binary_tree/main.go)

## Status

DONE while I though it was eazy

## Core idea

- DFS with a initial max value = math.MinInt so every value is bigger than it
- recursive, if the child value is bigger or equal than max, count++ and update max

## Failure

- I could right the correct approach from the start, think it was easy, then got beaten by some edge cases of my code, then overthinking it with 2 maxes, each for left and right
- But max is a int value, so when we pass into each recursion, it's copied by value, so the value doesn't get leaked into other branch

## Success

- Well at least I know the correct direction

## Tags

`tree` `dfs`
