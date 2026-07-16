# 684. Redundant Connection

## Link

[Leetcode](https://leetcode.com/problems/redundant-connection/description/)

## Difficulty

Medium

## Problem

In this problem, a tree is an undirected graph that is connected and has no cycles.

You are given a graph that started as a tree with n nodes labeled from 1 to n , with one additional edge added. The added edge has two different vertices chosen from 1 to n , and was not an edge that already existed. The graph is represented as an array edges of length n where edges[i] = [a i , b i ] indicates that there is an edge between nodes a i and b i in the graph.

Return an edge that can be removed so that the resulting graph is a tree of n nodes . If there are multiple answers, return the answer that occurs last in the input.

Example 1:

Input: edges = [[1,2],[1,3],[2,3]]
Output: [2,3]

Example 2:

Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
Output: [1,4]

Constraints:

n == edges.length

3

edges[i].length == 2

1 i i

a i != b i

There are no repeated edges.

The given graph is connected.

## Solve Date

20260716

## Solution

[Code](../../src/684_redundant_connection/main.go)

## Status

DONE

## Core idea

- use a technique call union-find, or disjoint set union
- create a parent int[] array, where index i is the node number i, and value of i is the index of node i's parent
- we initialize the array with each index set to itself, when no node are connected, then each node is its own parent
- we then create a find function, that trace the root of the target node, the root is found when its index == its value, meaning it doesn't point to any other node
- then we loop through the edges, we connect those 2 node by set one's root = the other's root using Find(), effectively connects them
- while looping, if any 2 node already shared the same root, meaning they already connected, new edge only create a loop, we return that two node

## Failure

- before this, i don't know what the fuck is union-join

## Success

- but now I know, haha

## Tags

`graph`
