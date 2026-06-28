# 152. Maximum Product Subarray

## Link

[Leetcode](https://leetcode.com/problems/maximum-product-subarray/description/)

## Difficulty

Medium

## Problem

Given an integer array nums , find a subarray that has the largest product, and return the product .

The test cases are generated so that the answer will fit in a 32-bit integer.

Note that the product of an array with a single element is the value of that element.

 Example 1: 

 Input: nums = [2,3,-2,4]
 Output: 6
 Explanation: [2,3] has the largest product 6.

 Example 2: 

 Input: nums = [-2,0,-1]
 Output: 0
 Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Constraints: 

 1 4 

-10 

The product of any subarray of nums is guaranteed to fit in a 32-bit integer.

## Solve Date

20260628

## Solution

[Code](../../src/152_maximum_product_subarray/main.go)

## Status

DONE

## Core idea

- a negative could negate the whole result from smallest to greatest in 1 calculation, so we need to keep track of the previous max and min
- there 3 possible product: maxProd[i-1] * i, minProd[i-1] * i; start fresh from i, we get the max and min in 3 of those
- one last loop to get the max in maxProd

## Failure

- still need hint

## Success

- new dp pattern, why so many?

## Tags

`1d-dp` `array`
