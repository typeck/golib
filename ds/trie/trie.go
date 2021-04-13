package trie

type Trie struct {
	child [128]*Trie
	exist bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	node := t
	for i := range word {
		if node.child[word[i]] == nil {
			c := Constructor()
			node.child[word[i]] = &c
		}
		node = node.child[word[i]]
	}
	node.exist = true
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.exist
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t.SearchPrefix(prefix)
	return node != nil
}

func (t *Trie)SearchPrefix(prefix string) *Trie {
	var node = t
	for i := range prefix {
		if node.child[prefix[i]] == nil {
			return nil
		}
		node = node.child[prefix[i]]
	}
	return node
}
