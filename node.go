package trie

import "sort"

type node struct {
	value     rune
	parent    *node
	children  []*node
	endOfWord bool
}

// NewNode initializes a node with the value
func newNode(r rune, suffix []rune, parent *node) (*node, bool) {

	var inserted bool

	n := &node{
		value:    r,
		children: make([]*node, 0),
		parent:   parent,
	}

	if len(suffix) > 0 {
		n.children, inserted = insert(n.children, suffix, n)
	} else {
		n.endOfWord = true
	}

	return n, inserted
}

func insert(nodes []*node, word []rune, parent *node) ([]*node, bool) {

	var inserted bool

	prefix, suffix := word[0], word[1:]

	index, n := search(nodes, prefix)
	if n != nil {
		if len(suffix) > 0 {
			n.children, inserted = insert(n.children, suffix, n)
		} else {
			inserted, n.endOfWord = true, true
		}
	} else {

		nodeToInsert, _ := newNode(prefix, suffix, parent)

		if index == len(nodes) {
			// If the new node should be on the end, just append it
			nodes = append(nodes, nodeToInsert)
		} else {
			// Otherwise insert the node at the appropriate index
			nodes = append(nodes, nil)
			copy(nodes[index+1:], nodes[index:])
			nodes[index] = nodeToInsert
		}

		inserted = true
	}

	return nodes, inserted
}

// contains recursively searches children until it reaches the end of the word,
// or it cannot find a node matching the rune
func contains(nodes []*node, word []rune) bool {

	// Search for the node that matches the rune
	_, node := search(nodes, word[0])
	if node == nil {
		return false
	}

	// if this is the last rune in the word, we don't search anymore children
	if endOfWord := len(word) == 1; endOfWord {
		return node.endOfWord
	}

	// recursively search the children
	return contains(node.children, word[1:])
}

// search looks for the node where the value matches the rune
func search(nodes []*node, r rune) (int, *node) {
	index := sort.Search(len(nodes), func(i int) bool { return nodes[i].value >= r })
	if index >= 0 && index < len(nodes) && nodes[index].value == r {
		return index, nodes[index]
	}

	return index, nil
}
