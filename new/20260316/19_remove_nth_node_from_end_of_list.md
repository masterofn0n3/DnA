# 19. Remove Nth Node From End of List

## Link

[Leetcode](https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/)

## Solve Date

20260318

## Solution

[Code](../../src/19_remove_nth_node_from_end_of_list/main.go)

## Status

DONE

## Core idea

use the same slow fast pattern, but the fast pointer is n position(n provided by params) ahead of slow
when the fast pointer reach the end, meaning the slow pointer is right before at the position we need to cut
- O(n)


## Failure

- need detail sample list to detail out the solution but I'm not sure if this is a failure, but I'm a self-loadthing ass so I'm gonna put it here
## Success

- Manage to solve it without too much hassle, nice

## Tags

`linked-list`
