package awesome


type Trie struct {
	children [26]*Trie
	isEnd bool
}


func Trie_Constructor() Trie {
	return Trie{}
}



func (this *Trie) Insert(word string)  {
	var insert  func(node *Trie, word string)
	insert = func(node *Trie, word string) {
		pos := word[0] - 'a'
		val := node.children[pos]
		// 没有才创建
		if val == nil {
			node.children[pos] = &Trie{}
		}
		if len(word) == 1 {
			node.children[pos].isEnd = true
			return
		}
		// 递归
		insert(node.children[pos], word[1:])
	}
	insert(this, word)
}


func (this *Trie) Search(word string) bool {

	if len(word) == 0 {
		return true
	}
	var search func(node *Trie, word string) bool
	search = func(node *Trie, word string) bool {

		pos := word[0] - 'a'
		c := node.children[pos]
		if c == nil {
			return false
		}
		if len(word) == 1 {
			return c.isEnd
		}
		return search(c, word[1:])
	}
	return search(this, word)
}


func (this *Trie) StartsWith(prefix string) bool {
	var search func(node *Trie, prefix string) bool
	search = func(node *Trie, prefix string) bool {
		pos := prefix[0] - 'a'
		c := node.children[pos]
		if c == nil {
			return false
		}
		if len(prefix) == 1 {
			return true
		}
		return search(c, prefix[1:])
	}
	return search(this, prefix)
}