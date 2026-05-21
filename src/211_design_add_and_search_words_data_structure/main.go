package main

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	current := this
	for _, v := range word {
		index := v - 'a'
		if current.children[index] == nil {
			current.children[index] = &WordDictionary{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(node *WordDictionary, remaining string) bool
	dfs = func(node *WordDictionary, remaining string) bool {
		if len(remaining) == 0 {
			return node.isEnd
		}
		current := remaining[0]
		left := remaining[1:]

		if current != '.' {
			child := node.children[current-'a']
			if child == nil {
				return false
			}
			return dfs(child, left)
		} else {
			for _, v := range node.children {
				if v != nil {
					if dfs(v, left) {
						return true
					}
				}
			}
			return false
		}
	}
	return dfs(this, word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
