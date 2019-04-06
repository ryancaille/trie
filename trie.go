package trie

// Trie is an implementation of a prefix tree used for storing and searching strings
type Trie interface {
	Count() int
	Insert(word string)
	Contains(word string) bool
}

type trie struct {
	count    int
	children []*node
}

// NewTrie initializes the Trie
func NewTrie() Trie {
	return &trie{}
}

func (t *trie) Count() int {
	return t.count
}

func (t *trie) Insert(word string) {

	// If word has a zero length, do nothing
	if len(word) == 0 {
		return
	}

	if c, inserted := insert(t.children, word, nil); inserted {
		t.children = c
		t.count++
	}
}

func (t *trie) Contains(word string) bool {
	return true
}
