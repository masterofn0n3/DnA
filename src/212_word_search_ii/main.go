package main

type Trie struct {
	isEnd    bool
	children [26]*Trie
	word     string
}

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	root := &Trie{}
	//convert words to trie
	for _, word := range words {
		current := root
		for _, v := range word {
			index := v - 'a'
			if current.children[index] == nil {
				current.children[index] = &Trie{}
			}
			current = current.children[index]
		}
		current.isEnd = true
		current.word = word
	}

	var dfs func(row, col int, node *Trie)

	dfs = func(row, col int, node *Trie) {
		//out of bound
		if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {
			return
		}
		curChar := board[row][col]
		// invalid path, not exist or visited
		if curChar == '#' || node.children[curChar-'a'] == nil {
			return
		}

		node = node.children[curChar-'a']
		if node.isEnd {
			result = append(result, node.word)
			node.isEnd = false
		}

		board[row][col] = '#'

		dfs(row-1, col, node)
		dfs(row+1, col, node)
		dfs(row, col+1, node)
		dfs(row, col-1, node)

		board[row][col] = curChar

	}

	for row := range board {
		for col := range board[row] {
			dfs(row, col, root)
		}
	}

	return result

}
