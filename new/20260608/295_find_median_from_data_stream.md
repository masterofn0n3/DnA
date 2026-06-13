# 295. Find Median from Data Stream

## Link

[Leetcode](https://leetcode.com/problems/find-median-from-data-stream/description/)

## Difficulty

Hard

## Problem

The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

For example, for arr = [2,3,4] , the median is 3 .

For example, for arr = [2,3] , the median is (2 + 3) / 2 = 2.5 .

Implement the MedianFinder class:

MedianFinder() initializes the MedianFinder object.

void addNum(int num) adds the integer num from the data stream to the data structure.

double findMedian() returns the median of all elements so far. Answers within 10 -5 of the actual answer will be accepted.

Example 1:

Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1); // arr = [1]
medianFinder.addNum(2); // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3); // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0

Constraints:

-10 5 5

There will be at least one element in the data structure before calling findMedian .

At most 5 \* 10 4 calls will be made to addNum and findMedian .

Follow up:

If all integer numbers from the stream are in the range [0, 100] , how would you optimize your solution?

If 99% of all integer numbers from the stream are in the range [0, 100] , how would you optimize your solution?

## Solve Date

20260613

## Solution

[Code](../../src/295_find_median_from_data_stream/main.go)

## Status

DONE in 30 mins

## Core idea

- 2 heap, 1 minheap, 1 maxheap, max heap to store the lower half of the array, min heap to store the upper
- when add a element, compare with root of max heap to decide which heap to push to
- when heap size > 2, rebalance the heap
- because each heap contain half, so the median is one of the root of the heaps, depends on the heaps size, if size is equal then (root1+root2)/2

## Failure

- my initial direction failed from beginning, the heap doesn't stored array in order

## Success

- after getting hints, I implement without error, cool

## Tags

`heap`
