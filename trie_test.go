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

func TestInsertWord(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foobar")
	if trie.Count() != 1 {
		t.Error("trie should have one word")
	}

	rootChildren := trie.RootChildren()
	if len(rootChildren) != 1 {
		t.Error("trie should have one root child")
	}

	if rootChildren[0].Value() != 'f' {
		t.Error("the first child should have a value of 'f'")
	}
}
