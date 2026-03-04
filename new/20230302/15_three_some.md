# 167 Two Sum II

## Link

[Leetcode](https://leetcode.com/problems/3sum/description/)

## Solution

[Code](../../src/15_three_some/main.go)

## Status

DONE

## Core idea

- Similar to 2 sum, just sort the array and convert one of the triplet to the _target_(by using a outer loop and take the current index), then it will be back to 2 sum problem
- O(n\*2)

## Failure

- probably get some rest first, walk or sth before dive straight to a new problem after just solve one, its kinda exhause
- pay attention to the edge cases of advancing the 2 pointer in the inner loop so it would out of bound
- still need tip for remove duplicates, and right now still kinda vague idea when I'm writing this line

## Tags

`array`, `two pointer`
