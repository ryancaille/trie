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

	r := word[0]
	suffix := word[1:]

	if len(nodes) == 0 {
		node, _ := newNode(r, suffix, parent)
		nodes = append(nodes, node)
		inserted = true
	} else {
		// TODO: create Search method to combine with search from contains
		index := sort.Search(len(nodes), func(i int) bool { return nodes[i].value >= r })
		if index >= 0 && index < len(nodes) && nodes[index].value == r {
			if len(suffix) > 0 {
				nodes[index].children, inserted = insert(nodes[index].children, suffix, nodes[index])
			} else {
				nodes[index].endOfWord = true
				inserted = true
			}
		}
	}

	return nodes, inserted
}

func contains(nodes []*node, word []rune) bool {

	r := word[0]
	endOfWord := len(word) == 1

	// TODO: create Search method to combine with search from insert
	index := sort.Search(len(nodes), func(i int) bool { return nodes[i].value >= r })
	if index >= 0 && index < len(nodes) && nodes[index].value == r {

		if endOfWord && nodes[index].endOfWord {
			return true
		}

		if endOfWord && !nodes[index].endOfWord {
			return false
		}

		return contains(nodes[index].children, word[1:])
	}

	return false
}
