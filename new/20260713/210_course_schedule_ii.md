# 210. Course Schedule II

## Link

[Leetcode](https://leetcode.com/problems/course-schedule-ii/description/)

## Difficulty

Medium

## Problem

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1 . You are given an array prerequisites where prerequisites[i] = [a i , b i ] indicates that you must take course b i first if you want to take course a i .

For example, the pair [0, 1] , indicates that to take course 0 you have to first take course 1 .

Return the ordering of courses you should take to finish all courses . If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array .

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

Example 3:

Input: numCourses = 1, prerequisites = []
Output: [0]

Constraints:

1

0

prerequisites[i].length == 2

0 i , b i

a i != b i

All the pairs [a i , b i ] are distinct .

## Solve Date

20260715

## Solution

[Code](../../src/210_course_schedule_ii/main.go)

## Status

DONE

## Core idea

- topological sort
- think of it this way, when using dfs, we construct the graph in direction, a needs b to complete, so a <- b, but to do bfs, it's when complete a, b will be available to learn, so a -> b
- we create a graph [][]int array to store that relationship, graph[1] = [2,3,4] means when complete course 1, we will be able to learn course 2,3,4 which is opposite to dfs where graph[1] = [2,3,4], means to learn course 1, we will need to learn 2,3,4 fisrt
- along while we construct the graph array, we also construct the inDegree array, in represent inDegree[1] = 3, means for course 1, we need 3 other course to start
- then do a loop through the inDegree array to push courses with degree 0 into a queue, those are the one that don't have prequisite
- for len(queue) > 0, we will pop the first element, push it to result, then decrease the degree of the courses it unlock, using the graph array, and if any course degree comes to 0, we again push it to the queue
- if len(result) == numCourse, that mean its possible to learn all courses, return result, otherwise return empty array
- O(V+E)

## Failure

## Success

- now I know topological sort muahahaah

## Tags

`graph
