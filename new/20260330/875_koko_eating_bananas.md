# 875. Koko Eating Bananas

## Link

[Leetcode](https://leetcode.com/problems/koko-eating-bananas/description/)

## Solve Date

20260330

## Solution

[Code](../../src/875_koko_eating_bananas/main.go)

## Status

DONE

## Core idea

- the max speed is equivalent to the largest pile
- min speed is 1 per hour
- use binary search to find the lowest speed possible
- ceil(a/b) = (a+b-1)/b

## Failure

- the ceil formular is still a blackbox to me, i haven't understand the math behind yet, so will forget if I don't review later
- still need hint
- i don't know the condition to move the left pointer forward

## Success

- implement it quite fast, since I have done this issue before

## Tags

`binary-search`
