Core Idea

- 2 pointer left and right
- mid = (left+right)/2
- base a given condition to increase left to mid + 1 or right to mid
- use this one when we see this pattern if x work, then x+1 work, .... or vice versa
- "find the minimum/maximum value in a range that satisfy the properties" or just to find a number in a range

- notes on condition
  - left <= right for when we want to find a exact value, left < right when we want to find a boundary
  - if mid is definitely invalid, skip it by left = mid + 1 or right = mid - 1, otherwise include it with right = mid or left = mid, usually it's both for each left and right
  - usually left < right pair with right = mid, and left = mid + 1
