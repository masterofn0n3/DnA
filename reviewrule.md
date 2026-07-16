Problem:
Find maximum of every sliding window.

Brute force:
Scan every window.

Observation:
Smaller elements behind a larger one are useless forever.

Data structure:
Monotonic decreasing deque.

Invariants:

1. Indices increasing.
2. Values decreasing.
