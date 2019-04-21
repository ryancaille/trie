package trie

import (
	"bufio"
	"log"
	"os"
	"sync"
	"testing"
)

var (
	wordsAlphabet        = readWordFile("test/alphabetical")
	wordsReverseAlphabet = readWordFile("test/reverseAlphabetical")
	wordsLike            = readWordFile("test/likeWords")
)

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

func TestInsertTheSameWordTwice(t *testing.T) {
	trie := NewTrie()
	trie.Insert("foobar")
	trie.Insert("foobar")
	if trie.Count() != 1 {
		t.Error("trie should have one word")
	}

	if !trie.Contains("foobar") {
		t.Error("trie should contain foobar")
	}
}

func TestInsertOverlappingWords(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foobar")
	trie.Insert("foo")

	if trie.Count() != 2 {
		t.Error("trie should have two words")
	}

	if !trie.Contains("foobar") {
		t.Error("trie should contain foobar")
	}

	if !trie.Contains("foo") {
		t.Error("trie should contain foo")
	}
}

func TestWordsInsertedInOrder(t *testing.T) {
	trie := NewTrie()

	for _, w := range wordsAlphabet {
		trie.Insert(w)
		if !trie.Contains(w) {
			t.Errorf("trie should contain %v", w)
		}
	}

	if trie.Count() != 26 {
		t.Error("trie should have 26 words")
	}
}

func TestWordsInsertedOutOfOrder(t *testing.T) {
	trie := NewTrie()

	for _, w := range wordsReverseAlphabet {
		trie.Insert(w)
		if !trie.Contains(w) {
			t.Errorf("trie should contain %v", w)
		}
	}

	if trie.Count() != 26 {
		t.Error("trie should have 26 words")
	}
}

func TestRemoveEmptyWordShouldDoNothing(t *testing.T) {
	trie := NewTrie()
	trie.Remove("")
}

func TestRemoveNonExistentWordShouldDoNothing(t *testing.T) {
	trie := NewTrie()
	trie.Remove("nothing")
}

func TestRemoveWord(t *testing.T) {

	trie := NewTrie()

	trie.Insert("remove")
	if trie.Count() != 1 {
		t.Error("trie should have one word")
	}

	if !trie.Contains("remove") {
		t.Error("trie should contain remove")
	}

	trie.Remove("remove")
	if trie.Count() != 0 {
		t.Error("trie should have zero words")
	}

	if trie.Contains("remove") {
		t.Error("trie should not contain remove")
	}

}

func TestRemoveOverlappingWord(t *testing.T) {

	trie := NewTrie()

	trie.Insert("foo")
	trie.Insert("foobar")
	trie.Remove("foobar")

	if !trie.Contains("foo") {
		t.Error("trie should contain foo")
	}

	if trie.Contains("foobar") {
		t.Error("trie should not contain foobar")
	}
}

func TestRemoveUnderlappingWord(t *testing.T) {

	trie := NewTrie()

	trie.Insert("foo")
	trie.Insert("foobar")
	trie.Remove("foo")

	if !trie.Contains("foobar") {
		t.Error("trie should contain foobar")
	}

	if trie.Contains("foo") {
		t.Error("trie should not contain foo")
	}
}

func TestConcurrentInserts(t *testing.T) {

	var wg sync.WaitGroup

	trie := NewTrie()
	words := []string{"aabc", "abbb", "abca", "aabb", "aaca", "accb", "acba", "bbac", "babc", "bbbb"}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, w := range words {
				trie.Insert(w)
			}
		}()
	}

	wg.Wait()

	if trie.Count() != len(words) {
		t.Errorf("Trie should contain %v words but found %v", len(words), trie.Count())
	}

	for _, w := range words {
		if !trie.Contains(w) {
			t.Errorf("Trie should contain %v and it does not", w)
		}
	}
}

func TestWordsLikeEmptyWord(t *testing.T) {

	trie := NewTrie()
	for _, w := range wordsLike {
		trie.Insert(w)
	}
	verifyMatches(t, trie.Like(""))
}

func TestWordsNoMatches(t *testing.T) {

	trie := NewTrie()
	for _, w := range wordsLike {
		trie.Insert(w)
	}
	verifyMatches(t, trie.Like("b"))
}

func TestWordsLikeWord(t *testing.T) {

	trie := NewTrie()
	for _, w := range wordsLike {
		trie.Insert(w)
	}

	verifyMatches(t, trie.Like("a"), "aachen", "aaron", "aaronite", "abaciscus", "abaco")
}

func verifyMatches(t *testing.T, actual []string, expected ...string) {
	if len(actual) != len(expected) {
		t.Fatalf("There should be %v matches but found %v", len(actual), len(expected))
	}
}

func readWordFile(file string) []string {
	words := make([]string, 0)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}
