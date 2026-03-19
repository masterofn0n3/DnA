# 287. Find the Duplicate Number

## Link

[Leetcode](https://leetcode.com/problems/find-the-duplicate-number/description/)

## Solve Date

20260319

## Solution

[Code](../../src/287_find_the_duplicate_number/main.go)

## Status

DONE with vague

## Core idea

- floyd cycle detection
  - consider the array is a linked list itself, where array values is pointer to next index
  - one slow and fast pointer starting at the same position, fast walk x2 slow, if there's repetition, they are guarantee to meet in a cycle
  - then move slow back to the start and walk both at the same speed until they're met again, then we have the repetition node
  - O(n)
- there's pigeonhole approach, but I haven't try yet

## Failure

- Still couldn't wrap my head around the math here, need to review back soon
- So I review right after because my brain need closure, and it's quite clear now
  - Until the slow and fast meet, slow walk F + a(F is length from the start to the cycle start, a is from the start to where they meet)
  - Fast have walk F + kC + a(C is cycle length, k is number of times fast have traverse through the cycle)
  - because fast walk fast x2 slow => F + a + kC = 2(F+a) => kC = F + a => F = kC - a
  - so if we move slow back to beginning and fast at the same point the met, slow walks a F path, fast will walk k time through the cycle, minus a, and they will meet at the start of the cycle => the value we need for this problem

## Success

<!-- What went well / what to reuse -->

## Tags

`linked-list` `high-priority-review`
