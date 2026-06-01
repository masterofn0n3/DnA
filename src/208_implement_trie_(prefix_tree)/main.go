package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, v := range word {
		index := v - 'a'
		if current.children[index] == nil {
			current.children[index] = &Trie{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for _, v := range word {
		index := v - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, v := range prefix {
		index := v - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
