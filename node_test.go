package trie

import (
	"fmt"
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
	node, inserted := newNode('f', []rune{'u', 'n'}, nil)

	if !inserted {
		t.Error("inserted should be true")
	}

	expectations := []nodeExpectation{
		{value: 'f', children: []rune{'u'}, nextChild: 'u'},
		{value: 'u', children: []rune{'n'}, nextChild: 'n', parent: 'f'},
		{value: 'n', parent: 'u', endOfWord: true},
	}

	verifyExpectations(t, node, expectations, 0, "f")
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

	if node.endOfWord {
		fmt.Printf("[%v] Validated Word\n", prefix)
	} else {
		fmt.Printf("[%v] Validated Prefix\n", prefix)
	}

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

	for _, r := range e.children {
		if !nodeHasChild(n, r) {
			t.Errorf("[%v] node should have a child '%c' and does not", prefix, r)
		}
	}

	for _, c := range n.children {
		if !isRuneExpected(e.children, c.value) {
			t.Errorf("[%v] node has a child '%c' but this child is not expected", prefix, c.value)
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

func nodeHasChild(n *node, r rune) bool {
	for _, c := range n.children {
		if c.value == r {
			return true
		}
	}

	return false
}

func isRuneExpected(expected []rune, r rune) bool {
	for _, e := range expected {
		if e == r {
			return true
		}
	}

	return false
}

func getNextNode(children []*node, r rune) *node {
	for _, c := range children {
		if c.value == r {
			return c
		}
	}

	return nil
}
