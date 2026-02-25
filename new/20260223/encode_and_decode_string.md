# 217. Encode and Decode String

## Link

[Neetcode](https://neetcode.io/problems/string-encode-and-decode/history?list=neetcode150&submissionIndex=0)

## Solution

[Code](src/encode_and_decode_string/main.go)

## Status

DONE in nearly an hour

## Core idea

Create a delimeter = a symbol + word length, then traverse the string using that delim

O(n)

## Failure

- Use string.Builder instead of result += because every string concat allo a new string
- Use string.Index to find index of a char instead of using a inner loop
- Can drop the first # to simplify the delim, also remove the infinite loop risks

## Tags

`array`, `map`
