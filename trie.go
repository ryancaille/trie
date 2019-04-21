package trie

import (
	"strings"
	"sync"
)

// Trie is a data structure that is optimized for storing and searching strings,
// as well as string matching based on a prefix
type Trie struct {
	count    int
	children []*node
	lock     sync.RWMutex
}

// NewTrie initializes the Trie
func NewTrie() *Trie {
	return &Trie{}
}

// Count returns the number of unique words currently stored in the Trie
func (t *Trie) Count() int {

	t.lock.RLock()
	c := t.count
	t.lock.RUnlock()

	return c
}

// Insert will insert a new word into the Trie.  If the word already exists it does not store a duplicate copy,
// nor will it track how many times that word was inserted, only that it does exist in the Trie.
func (t *Trie) Insert(word string) {

	// If word has a zero length, do nothing
	if len(word) == 0 {
		return
	}

	t.lock.Lock()
	if c, inserted := insert(t.children, splitWord(word), nil); inserted {
		t.children = c
		t.count++
	}
	t.lock.Unlock()
}

// Contains will check the Trie to see if a word is currently stored.
func (t *Trie) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}

	t.lock.RLock()
	found, _ := contains(t.children, splitWord(word))
	t.lock.RUnlock()

	return found
}

// Remove will remove a word if it exists, and do nothing if it does not exist
func (t *Trie) Remove(word string) {
	if len(word) == 0 {
		return
	}

	t.lock.Lock()
	if c, removed := remove(t.children, splitWord(word)); removed {
		t.children = c
		t.count--
	}
	t.lock.Unlock()
}

// Like will traverse the Trie and find the best matches. up to the supplied count
func (t *Trie) Like(prefix string, count int) []string {

	if len(prefix) == 0 {
		return make([]string, 0)
	}

	t.lock.RLock()
	words := like(t.children, splitWord(prefix), count)
	t.lock.RUnlock()

	return words
}

func splitWord(word string) []rune {
	return []rune(strings.ToLower(word))
}
