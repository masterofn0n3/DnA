# 49. Group Anagrams

## Link

[Leetcode](https://leetcode.com/problems/group-anagrams/submissions/1931404239/)

## Solution

[Code](src/49_group_anagram/main.go)

## Status

DONE in nearly an hour

## Core idea

- Bruceforce by loop and compare each index and group them together, remove the grouped element in the processs, make sure to decrease
  the index because we removed it, then the loop will ++
- Using map to indentify aragram by sorting them, or create a 26 char array as key,

O(n*2 *k)
O(n \* k log n)

## Failure

- draft out the bruceforce solution first, then figure out if I can optimize it, if not, just code the bruceforce
- Review how to manipulate slice: remove, add,... and how to create resource in go, forgot it because long time no use
- Need to understand log k and more deeply about time and space complexity

## Tags

`array`, `map`
