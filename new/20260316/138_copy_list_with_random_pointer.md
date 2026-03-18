# 138. Copy List with Random Pointer

## Link

[Leetcode](https://leetcode.com/problems/copy-list-with-random-pointer/description/)

## Solve Date

20260318

## Solution

[Code](../../src/138_copy_list_with_random_pointer/main.go)

## Status

DONE in the midnight

## Core idea

- use map to store the original node as key with new node as value with 1 pass
- the second pass to go through the list again and set the value using the map
- time O(n) space O(1) since we use additional map
- more optimize way is use interleaving approach
  - insert new node right next to original node
  - set new node.random using original node.random.next, because original node is right before it, same for the random original
  - separate the list again => nice approach

## Failure

- first time clone so I don't know how, but now I know, create a dummy node, then go through the list, set the dummy.Next to the current node, and prev node to the current node, so the dummy will point to the `head` clone
- Got to sleepy and tunnel vision to recognize the use of map, better to solve this in the day
- I could solve it in the evening, but I got too distracted with many distracting stuff

## Success

- could reuse the cloning and interleaving pattern(nice pattern)

## Tags

`linked-list` `map` `interleaving-approach`
