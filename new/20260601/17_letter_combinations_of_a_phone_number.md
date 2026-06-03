# 17. Letter Combinations of a Phone Number

## Link

[Leetcode](https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/)

## Difficulty

Medium

## Problem

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order .

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:

Input: digits = "2"
Output: ["a","b","c"]

Constraints:

1

digits[i] is a digit in the range ['2', '9'] .

## Solve Date

20260602

## Solution

[Code](../../src/17_letter_combinations_of_a_phone_number/main.go)

## Status

DONE

## Core idea

- same idea with [palindrom partitioning](131_palindrome_partitioning.md), but simpler
- a map to store the letter for each number
- a for loop to interate through each digit, then recursive call with index + 1 in each iteration

## Failure

- simple but i still need hint to get the right recursive direction, but i get the base right tho

## Success

## Tags

`backtracking` `recursion`
