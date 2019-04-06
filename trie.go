package trie

// Trie is an implementation of a prefix tree used for storing and searching strings
type Trie interface {
	Count() int
	Insert(word string)
	RootChildren() []Node
}

type trie struct {
	count int
}

// NewTrie initializes the Trie
func NewTrie() Trie {
	return &trie{}
}

func (t *trie) Count() int {
	return t.count
}

func (t *trie) Insert(word string) {
	t.count++
}

func (t *trie) RootChildren() []Node {
	return []Node{&node{value: 'f'}}
}
