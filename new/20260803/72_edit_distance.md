# 72. Edit Distance

## Link

[Leetcode](https://leetcode.com/problems/edit-distance/description/)

## Difficulty

Medium

## Problem

Given two strings word1 and word2 , return the minimum number of operations required to convert word1 to word2 .

You have the following three operations permitted on a word:

 Insert a character

Delete a character

Replace a character

 Example 1: 

 Input: word1 = "horse", word2 = "ros"
 Output: 3
 Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

 Example 2: 

 Input: word1 = "intention", word2 = "execution"
 Output: 5
 Explanation: 
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')

Constraints: 

 0 

word1 and word2 consist of lowercase English letters.

## Solve Date

20260807

## Solution

[Code](../../src/72_edit_distance/main.go)

## Status

DONE

## Core idea
- base case: if i == len(word1) => we need len(word2) - j insert operation, else if j == len(word2) we need len(word1) - i delete operation
- for index i and j
    - if char at i = char at j => increase both without needing any operation
    - else +1 operation count and take the min of
        - replace: increase both i and j
        - insert: keep i and increase j
        - delete: increase i and keep i

## Failure

- wow I only just done 1 problem this week, what a failure

## Success

- well at least I got 1 instead of 0 :D 

## Tags

`dp`
