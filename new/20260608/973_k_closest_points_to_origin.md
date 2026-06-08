# 973. K Closest Points to Origin

## Link

[Leetcode](https://leetcode.com/problems/k-closest-points-to-origin/description/)

## Difficulty

Medium

## Problem

Given an array of points where points[i] = [x i , y i ] represents a point on the X-Y plane and an integer k , return the k closest points to the origin (0, 0) .

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x 1 - x 2 ) 2 + (y 1 - y 2 ) 2 ).

You may return the answer in any order . The answer is guaranteed to be unique (except for the order that it is in).

 Example 1: 

 Input: points = [[1,3],[-2,2]], k = 1
 Output: [[-2,2]]
 Explanation: 
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) Example 2: 

 Input: points = [[3,3],[5,-1],[-2,4]], k = 2
 Output: [[3,3],[-2,4]]
 Explanation: The answer [[-2,4],[3,3]] would also be accepted.

Constraints: 

 1 4 

-10 4 i , y i 4

## Solve Date

20260608

## Solution

[Code](../../src/973_k_closest_points_to_origin/main.go)

## Status

DONE

## Core idea

- create a heap
- pop k times
- done => O(n + k log n)
- more optimize way, use a max heap, add k element to it, for remaining element, pop and push => O(n log k)

## Failure

<!-- What went wrong / what to avoid next time -->

## Success

- lol why this problem is medium
- its exactly the same as previous one

## Tags

<!-- e.g. array, hash-set, sorting -->
