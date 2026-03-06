# 42. Trapping Rain Water

## Link

[Leetcode](https://leetcode.com/problems/trapping-rain-water/description/)

## Solution

[Code](../../src/42_trapping_rain_water/main.go)

## Status

DONE

## Core idea

- calculate the water each column can hold, each column can hold the ower column of the 2 max outer column minus its own height
- bruce force with O(n\*2) by do 2 additional inner loop to find the left and right max for each index
- optimize by using 2 array to store left and right max of each index and one last loop to calculate the water volumn => O(n)
- optimize again by only using one array to store right max, left max is calculated and accumulate along when we go through the array again to calculate the total water

## Failure

- Last time 2 years ago I was able to figure this out myself, but took a lot of time, this time faster but I need tip, and can only figure out the bruceforce version, need another tip for the optimize version
- Next time review I need to do the optimize version with 2 loops

## Success

<!-- What went well / what to reuse -->

## Tags

`array` `2 pointers` `prefix product`
