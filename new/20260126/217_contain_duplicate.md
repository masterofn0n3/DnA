# 217. Contains Duplicate

## Link

[LeetCode](https://leetcode.com/problems/contains-duplicate/description/)

## Solution

[Code](../../src/contain_duplicate/main.go)

## Status

DONE

## Core idea

Loop through the array once, store the occurence in a map, return true if encounter one existing value, otherwise return false

O(n)

## Failure

Coulda use map with struct{}{} value for better memory saving, but instead I use a int - int map

## Tags

`array`, `map`