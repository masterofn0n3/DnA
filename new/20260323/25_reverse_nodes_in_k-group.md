# 25. Reverse Nodes in k-Group

## Link

[Leetcode](https://leetcode.com/problems/reverse-nodes-in-k-group/description/)

## Solve Date

20260329

## Solution

[Code](../../src/25_reverse_nodes_in_k-group/main.go)

## Status

DONE

## Core idea

- just a bunch of reassigning nodes
- one outer loop to traverse the nodes
- two inner loop, for checking if the number of nodes enough k nodes to reverse, another one to actually reverse the nodes in the group
- keep track of the groupTail to update the pointers after the loop finish

## Failure

- too many variable that need to be kept track, my working memory is overload, need a white board to visualize
- next time start with a small set of nodes to figure out the logic first, it usually the same logic for arbitrary nodes number
- took like 3 hour of incontinuous thinking, its 1.59PM now
- also need suggestion from AI instructor
- O(n)

## Success

<!-- What went well / what to reuse -->

## Tags

`linked-list`

