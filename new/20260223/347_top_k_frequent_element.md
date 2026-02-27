# 347. Top K Frequent Elements

## Link

[Leetcode](https://leetcode.com/problems/top-k-frequent-elements/description/)

## Solution

[Code](src/347_top_k_frequent_element/main.go)

## Status

DONE in nearly an hour

## Core idea

- use a map to store frequency of each element
- then use a array with frequency as index

O(n log n)

## Failure

- each time notice `frequecy` or `top K element`, think array index/ bucket, store frequency as an array index

## Tags

`array`, `map`
