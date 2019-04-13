package trie

import "strings"

// Trie is a data structure that is optimized for storing and searching strings,
// as well as string matching based on a prefix
type Trie struct {
	count    int
	children []*node
}

// NewTrie initializes the Trie
func NewTrie() *Trie {
	return &Trie{}
}

// Count returns the number of unique words currently stored in the Trie
func (t *Trie) Count() int {
	return t.count
}

// Insert will insert a new word into the Trie.  If the word already exists it does not store a duplicate copy,
// nor will it track how many times that word was inserted, only that it does exist in the Trie.
func (t *Trie) Insert(word string) {

	// If word has a zero length, do nothing
	if len(word) == 0 {
		return
	}

	if c, inserted := insert(t.children, splitWord(word), nil); inserted {
		t.children = c
		t.count++
	}
}

// Contains will check the Trie to see if a word is currently stored.
func (t *Trie) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}

	return contains(t.children, splitWord(word))
}

func splitWord(word string) []rune {
	return []rune(strings.ToLower(word))
}
