package trie

import "testing"

func TestCreate(t *testing.T) {
	trie := NewTrie()

	if trie == nil {
		t.Error("trie cannot be nil")
	}

	if trie.Count() != 0 {
		t.Error("trie should have zero words after creation")
	}
}

func TestInsertEmptyWord(t *testing.T) {
	trie := NewTrie()

	trie.Insert("")

	if trie.Count() != 0 {
		t.Error("trie should not insert empty word")
	}
}

func TestInsertWord(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foobar")
	if trie.Count() != 1 {
		t.Error("trie should have one word")
	}

	if !trie.Contains("foobar") {
		t.Error("trie should contain foobar")
	}
}
