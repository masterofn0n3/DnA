# 703. Kth Largest Element in a Stream

## Link

[Leetcode](https://leetcode.com/problems/kth-largest-element-in-a-stream/description/)

## Difficulty

Easy

## Problem

You are part of a university admissions office and need to keep track of the kth highest test score from applicants in real-time. This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.

You are tasked to implement a class which, for a given integer k , maintains a stream of test scores and continuously returns the k th highest test score after a new score has been submitted. More specifically, we are looking for the k th highest score in the sorted list of all scores.

Implement the KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums .

int add(int val) Adds a new test score val to the stream and returns the element representing the k th largest element in the pool of test scores so far.

Example 1:

Input:

["KthLargest", "add", "add", "add", "add", "add"]

[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

Output: [null, 4, 5, 5, 8, 8]

Explanation:

KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);

kthLargest.add(3); // return 4

kthLargest.add(5); // return 5

kthLargest.add(10); // return 5

kthLargest.add(9); // return 8

kthLargest.add(4); // return 8

Example 2:

Input:

["KthLargest", "add", "add", "add", "add"]

[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

Output: [null, 7, 7, 7, 8]

Explanation:

KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);

kthLargest.add(2); // return 7

kthLargest.add(10); // return 7

kthLargest.add(9); // return 7

kthLargest.add(9); // return 8

Constraints:

0 4

1

-10 4 4

-10 4 4

At most 10 4 calls will be made to add .

## Solve Date

20260605

## Solution

[Code](../../src/703_kth_largest_element_in_a_stream/main.go)

## Status

DONE in one morning

## Core idea

- simple way: sort the array, and when adding, add it to correct position, then use arr[k-1] to get kth element
- cool way:
  - use a [heap](../../pattern/heap.md), in this case `min heap`
  - when creating the struct, initialize the heap, start from the last parent node of the heap to bubbling up values, after finish creating the heap, pop the root and bubbeling down(sift down) until only K element in the heap
    - for popping, switch value of the root and the end of array, then remove the last element, or the heap will have a hole in it
  - when adding a new value, sift up the value, then pop until only k element left in the heap

## Failure

- first time I encounter this type of dsa, no idea how to do

## Success

- but i manage to do the easy approach tho

## Tags

`heap`
