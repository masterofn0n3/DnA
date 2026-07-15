# 207. Course Schedule

## Link

[Leetcode](https://leetcode.com/problems/course-schedule/description/)

## Difficulty

Medium

## Problem

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1 . You are given an array prerequisites where prerequisites[i] = [a i , b i ] indicates that you must take course b i first if you want to take course a i .

For example, the pair [0, 1] , indicates that to take course 0 you have to first take course 1 .

Return true if you can finish all courses. Otherwise, return false .

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.

Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

Constraints:

1

0

prerequisites[i].length == 2

0 i , b i

All the pairs prerequisites[i] are unique .

## Solve Date

20260714

## Solution

[Code](../../src/207_course_schedule/main.go)

## Status

DONE

## Core idea

- main idea: if the graph contain a loop, that means it couldn't be complete
- dfs way
  - first we need to convert the course and its prequisite into a graph representative, aka a `[][]int`, each index contain the current course's prequisite course
  - then traverse it in dfs manner, each node could have 3 status, unvisited, visited, done, which can be represent by a []int array with number 0 1 2
  - in dfs, if we encounter visited node, that mean we traverse back to the node in the visiting path => a loop => return false immediately, otherwise return true
  - after the traversal is finished, set the node status to done, so other will skip it, no necessary processing

## Failure

- i need correction on the idea, and also direction on how to implement the status track

## Success

- deduce and implement most of these myself, and only 2 point need to fix after, awesome!!

## Tags

`graph`
