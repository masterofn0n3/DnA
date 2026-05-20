# 208. Implement Trie (Prefix Tree)

## Link

[Leetcode](https://leetcode.com/problems/implement-trie-prefix-tree/description/)

## Difficulty

Medium

## Problem

A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.

void insert(String word) Inserts the string word into the trie.

boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.

boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix , and false otherwise.

Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple"); // return True
trie.search("app"); // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app"); // return True

Constraints:

1

word and prefix consist only of lowercase English letters.

At most 3 \* 10 4 calls in total will be made to insert , search , and startsWith .

## Solve Date

20260520

## Solution

[Code](<../../src/208_implement_trie_(prefix_tree)/main.go>)

## Status

DONE in 1 distracted hour, so headache

## Core idea

- new concept of trie, or prefix tree, a tree has nodes and each node can have many children
- a node is a character, each has 26 children, represent 26 lower case char, so every path represent a prefix of some word
- for loop the word to insert and search it in the tree
- O(n) with n is word length

## Failure

- need hints, but new concept

## Success

- i learned new concept

## Tags

`trie`
