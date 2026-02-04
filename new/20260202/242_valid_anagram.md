<!-- Copy to new/YYYYMMDD/NUMBER_slug.md, then replace placeholders below -->

## Link

[LeetCode](https://leetcode.com/problems/valid-anagram/)

## Solution

[Code](../../src/242_valid_anagram/main.go)

## Status

DONE in 38 mins

## Core idea

1. **Char array as frequency map** — length 26 for 'a'–'z'.
2. **Two loops:**
   - First: traverse both strings — increment count for each char in `s`, decrement for each char in `t`.
   - Second: check the 26 slots; if any count ≠ 0 → not an anagram.
3. **Index from rune:** for lowercase English, `rune - 'a'` gives index 0–25 (e.g. `'a' - 'a'` ⇒ 0).
4. **Time:** O(n). Optional: early return if `len(s) != len(t)`.

## Failure

- First attempt: solved for "is reverse?" instead of anagram → read the description carefully.
- Only 26 lowercase letters → use fixed array instead of map for better memory/performance → check input constraints.

## Tags

`array` `string` `counting`
