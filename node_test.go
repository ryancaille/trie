package trie

import (
	"testing"
)

type nodeExpectation struct {
	value     rune
	children  []rune
	nextChild rune
	parent    rune
	endOfWord bool
}

func TestSingleWordShouldBeInserted(t *testing.T) {

	root := make([]*node, 0)
	root = insertWordAndVerify(t, root, "fun", 1)

	actual := root[0]
	expectations := []nodeExpectation{
		{value: 'f', children: []rune{'u'}, nextChild: 'u'},
		{value: 'u', children: []rune{'n'}, nextChild: 'n', parent: 'f'},
		{value: 'n', parent: 'u', endOfWord: true},
	}

	verifyExpectations(t, actual, expectations, 0, "f")
}

func TestOverlappingWordsShouldBeInserted(t *testing.T) {

	root := make([]*node, 0)
	root = insertWordAndVerify(t, root, "fun", 1)
	root = insertWordAndVerify(t, root, "funny", 1)

	actual := root[0]
	expectations := []nodeExpectation{
		{value: 'f', children: []rune{'u'}, nextChild: 'u'},
		{value: 'u', children: []rune{'n'}, nextChild: 'n', parent: 'f'},
		{value: 'n', children: []rune{'n'}, nextChild: 'n', parent: 'u', endOfWord: true},
		{value: 'n', children: []rune{'y'}, nextChild: 'y', parent: 'n'},
		{value: 'y', parent: 'n', endOfWord: true},
	}

	verifyExpectations(t, actual, expectations, 0, "f")
}

func TestNodesWhenWordsAreInsertedOutOfOrder(t *testing.T) {

	root := make([]*node, 0)
	root = insertWordAndVerify(t, root, "gamma", 1)
	root = insertWordAndVerify(t, root, "beta", 2)
	root = insertWordAndVerify(t, root, "alpha", 3)

	alpha := root[0]
	beta := root[1]
	gamma := root[2]

	alphaExpectations := []nodeExpectation{
		{value: 'a', children: []rune{'l'}, nextChild: 'l'},
		{value: 'l', children: []rune{'p'}, nextChild: 'p', parent: 'a'},
		{value: 'p', children: []rune{'h'}, nextChild: 'h', parent: 'l'},
		{value: 'h', children: []rune{'a'}, nextChild: 'a', parent: 'p'},
		{value: 'a', parent: 'h', endOfWord: true}}

	betaExpectations := []nodeExpectation{
		{value: 'b', children: []rune{'e'}, nextChild: 'e'},
		{value: 'e', children: []rune{'t'}, nextChild: 't', parent: 'b'},
		{value: 't', children: []rune{'a'}, nextChild: 'a', parent: 'e'},
		{value: 'a', parent: 't', endOfWord: true}}

	gammaExpectations := []nodeExpectation{
		{value: 'g', children: []rune{'a'}, nextChild: 'a'},
		{value: 'a', children: []rune{'m'}, nextChild: 'm', parent: 'g'},
		{value: 'm', children: []rune{'m'}, nextChild: 'm', parent: 'a'},
		{value: 'm', children: []rune{'a'}, nextChild: 'a', parent: 'm'},
		{value: 'a', parent: 'm', endOfWord: true}}

	verifyExpectations(t, alpha, alphaExpectations, 0, "a")
	verifyExpectations(t, beta, betaExpectations, 0, "b")
	verifyExpectations(t, gamma, gammaExpectations, 0, "g")
}

func TestNodesWhenInsertingForkedWords(t *testing.T) {

	root := make([]*node, 0)
	root = insertWordAndVerify(t, root, "crazy", 1)
	root = insertWordAndVerify(t, root, "crayon", 1)
	root = insertWordAndVerify(t, root, "cream", 1)

	actual := root[0]

	crazyExpectations := []nodeExpectation{
		{value: 'c', children: []rune{'r'}, nextChild: 'r'},
		{value: 'r', children: []rune{'a', 'e'}, nextChild: 'a', parent: 'c'},
		{value: 'a', children: []rune{'y', 'z'}, nextChild: 'z', parent: 'r'},
		{value: 'z', children: []rune{'y'}, nextChild: 'y', parent: 'a'},
		{value: 'y', parent: 'z', endOfWord: true}}

	crayonExpectations := []nodeExpectation{
		{value: 'c', children: []rune{'r'}, nextChild: 'r'},
		{value: 'r', children: []rune{'a', 'e'}, nextChild: 'a', parent: 'c'},
		{value: 'a', children: []rune{'y', 'z'}, nextChild: 'y', parent: 'r'},
		{value: 'y', children: []rune{'o'}, nextChild: 'o', parent: 'a'},
		{value: 'o', children: []rune{'n'}, nextChild: 'n', parent: 'y'},
		{value: 'n', parent: 'o', endOfWord: true}}

	creamExpectations := []nodeExpectation{
		{value: 'c', children: []rune{'r'}, nextChild: 'r'},
		{value: 'r', children: []rune{'a', 'e'}, nextChild: 'e', parent: 'c'},
		{value: 'e', children: []rune{'a'}, nextChild: 'a', parent: 'r'},
		{value: 'a', children: []rune{'m'}, nextChild: 'm', parent: 'e'},
		{value: 'm', parent: 'a', endOfWord: true}}

	verifyExpectations(t, actual, crazyExpectations, 0, "c")
	verifyExpectations(t, actual, crayonExpectations, 0, "c")
	verifyExpectations(t, actual, creamExpectations, 0, "c")
}

func TestNodesAreRemovedWhenWordIsRemoved(t *testing.T) {
	root := make([]*node, 0)
	root = insertWordAndVerify(t, root, "remove", 1)
	root = removeWordAndVerify(t, root, "remove", 0)
}

func TestNodesAreCorrectAfterRemovingOverlappingWord(t *testing.T) {

	root := make([]*node, 0)
	root = insertWordAndVerify(t, root, "ab", 1)
	root = insertWordAndVerify(t, root, "abc", 1)
	root = removeWordAndVerify(t, root, "abc", 1)

	expectations := []nodeExpectation{
		{value: 'a', children: []rune{'b'}, nextChild: 'b'},
		{value: 'b', parent: 'a', endOfWord: true}}

	verifyExpectations(t, root[0], expectations, 0, "a")
}

func TestNodesAreCorrectAfterRemovingUnderlappingWord(t *testing.T) {

	root := make([]*node, 0)
	root = insertWordAndVerify(t, root, "ab", 1)
	root = insertWordAndVerify(t, root, "abc", 1)
	root = removeWordAndVerify(t, root, "ab", 1)

	expectations := []nodeExpectation{
		{value: 'a', children: []rune{'b'}, nextChild: 'b'},
		{value: 'b', children: []rune{'c'}, nextChild: 'c', parent: 'a'},
		{value: 'c', parent: 'b', endOfWord: true}}

	verifyExpectations(t, root[0], expectations, 0, "a")
}

func insertWordAndVerify(t *testing.T, root []*node, word string, expectedLen int) []*node {
	var inserted bool

	root, inserted = insert(root, []rune(word), nil)

	if !inserted {
		t.Errorf("%v should have been inserted and was not", word)
	}

	if len(root) != expectedLen {
		t.Errorf("root nodes should contain %v; found %v", expectedLen, len(root))
	}

	return root
}

func removeWordAndVerify(t *testing.T, root []*node, word string, expectedLen int) []*node {
	var removed bool

	root, removed = remove(root, []rune(word))

	if !removed {
		t.Fatalf("%v should have been removed and was not", word)
	}

	if len(root) != expectedLen {
		t.Fatalf("root nodes should contain %v; found %v", expectedLen, len(root))
	}

	return root
}

func verifyExpectations(t *testing.T, node *node, expectations []nodeExpectation, index int, prefix string) {
	if index >= len(expectations) {
		return
	}

	expect := expectations[index]

	validateValue(t, expect, node, prefix)
	validateParent(t, expect, node, prefix)
	validateChildren(t, expect, node, prefix)
	validateEndOfWord(t, expect, node, prefix)

	if expect.nextChild != 0 {

		if nextNode := getNextNode(node.children, expect.nextChild); nextNode != nil {
			index++
			verifyExpectations(t, nextNode, expectations, index, prefix+string(nextNode.value))
		} else {
			t.Errorf("[%v] cannot verify next node '%c' because it does not exist", prefix, expect.nextChild)
		}

	}
}

func validateValue(t *testing.T, e nodeExpectation, n *node, prefix string) {

	if value := n.value; value != e.value {
		t.Errorf("[%v] value must be '%c'; found %c", prefix, e.value, value)
	}
}

func validateParent(t *testing.T, e nodeExpectation, n *node, prefix string) {

	if e.parent != 0 && n.parent.value != e.parent {
		t.Errorf("[%v] node should have parent '%c', found '%c'", prefix, e.parent, n.parent.value)
	}

	if e.parent == 0 && n.parent != nil {
		t.Errorf("[%v] node should not have parent", prefix)
	}
}

func validateChildren(t *testing.T, e nodeExpectation, n *node, prefix string) {

	if len(e.children) != len(n.children) {
		t.Errorf("[%v] node should have ", prefix)
	}

	for i, r := range e.children {
		if n.children[i].value != r {
			t.Errorf("[%v] node should have a child '%c' at index %v and does not", prefix, r, i)
		}
	}

	for i, c := range n.children {
		if e.children[i] != c.value {
			t.Errorf("[%v] node has a child '%c' at index %v but this child is not expected", prefix, c.value, i)
		}
	}
}

func validateEndOfWord(t *testing.T, e nodeExpectation, n *node, prefix string) {

	if e.endOfWord && !n.endOfWord {
		t.Errorf("[%v] node should be end of word", prefix)
	}

	if !e.endOfWord && n.endOfWord {
		t.Errorf("[%v] node should NOT be end of word", prefix)
	}
}

func getNextNode(children []*node, r rune) *node {
	for _, c := range children {
		if c.value == r {
			return c
		}
	}

	return nil
}
