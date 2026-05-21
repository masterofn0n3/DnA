# 211. Design Add and Search Words Data Structure

## Link

[Leetcode](https://leetcode.com/problems/design-add-and-search-words-data-structure/description/)

## Difficulty

Medium

## Problem

Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

WordDictionary() Initializes the object.

void addWord(word) Adds word to the data structure, it can be matched later.

bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.

Example:

Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True

Constraints:

1

word in addWord consists of lowercase English letters.

word in search consist of '.' or lowercase English letters.

There will be at most 2 dots in word for search queries.

At most 10 4 calls will be made to addWord and search .

## Solve Date

20260521

## Solution

[Code](../../src/211_design_add_and_search_words_data_structure/main.go)

## Status

DONE

## Core idea

- similar to the 208 problem, prefix tree, addWord is exactly the same
- for search, because `.` match every character, we need a recursive function to check every child if a character is '.', if not `.` then call directly on the children[index]

## Failure

- still need hint
- i still need to improve on implementing recursion

## Success

- quite smoothly solve

## Tags

`trie` `recursion`
