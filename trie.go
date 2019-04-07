package trie

import "strings"

// Trie provides methods for storing and searching strings
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

	if c, inserted := insert(t.children, splitWord(word), nil); inserted {
		t.children = c
		t.count++
	}
}

func (t *trie) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}

	return contains(t.children, splitWord(word))
}

func splitWord(word string) []rune {
	return []rune(strings.ToLower(word))
}
