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

func TestShouldNotContainEmptyWord(t *testing.T) {
	trie := NewTrie()

	if trie.Contains("") {
		t.Error("trie should not contain empty word")
	}
}

func TestShouldNotContainWordWhenEmpty(t *testing.T) {
	trie := NewTrie()

	if trie.Contains("whatever") {
		t.Error("trie should not contain \"whatever\"")
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

	if trie.Contains("foo") {
		t.Error("trie should not contain foo")
	}
}
