// Package trie implements a search tree that stores strings that can be searched
// by a prefix. This would generally be used for an autocomplete feature where a
// input would return likely matches.
//
// Both reads and write are thread safe; however, one one write may occur at any
// one time.  Concurrent reads are not constrained.
//
// Usage would start by creating the Trie
//
//	trie := NewTrie()
//
// Autocomplete
//
// The main use case is for autocomplete. The Trie contains potential words, and
// a prefix is supplied to the Like() method.  All the words that start with the
// prefix are returned as a slice. You may specify the number of words you want
// returned.
//
//	words := trie.Like("foo", 5)
//
// CRUD Operations
//
// The Trie starts empty, so you would need to populate it to get any value back.
//
// You may add new strings to the Trie by invoking Insert(). Subsequent inserts
// of the same string to the Trie will not increase the memory space, only the
// existence of the word will be stored.
//
//	trie.Insert("foobar")
//
// You may check whether the word exists in the Trie by invoking Contains().
//
//	if trie.Contains("foobar") {
//		fmt.Print("found foobar")
//	}
//
// Trie stores all strings as lowercase. Asking whether a Trie contains "foobar"
// or "FOOBAR" would yield the same results.
//
// You may see how many unique words are stored in the Trie by invoking Count().
//
//	count := trie.Count()
//
// If you need to remove a word from the Trie, invoke Remove().
//
//	trie.Remove("foobar")
//
package trie
