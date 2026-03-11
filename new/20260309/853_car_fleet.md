# 853. Car Fleet

## Link

[Leetcode](https://leetcode.com/problems/car-fleet/description/)

## Solve Date

20260311

## Solution

[Code](../../src/853_car_fleet/main.go)

## Status

DONE

## Core idea

- sort the position in ascending order to see if which car is closest to the destination
- calculate the time of each car to reach the destination
- so the car with less time could reach the car with greater time and form a fleet, car with greater time form a fleet of its own(could use stack in this step to get only separate times)
- O(n logn) because of sorting

## Failure

- still need a lot of hints to solve this
- not focus
- I still not truly understand the idea of this issue

## Success

<!-- What went well / what to reuse -->

## Tags

`array` `stack` `high-priority-for-review`
