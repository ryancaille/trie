package trie

type node struct {
	value     rune
	parent    *node
	children  []*node
	endOfWord bool
}

// NewNode initializes a node with the value
func newNode(r rune, suffix string, parent *node) (*node, bool) {

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

func insert(nodes []*node, word string, parent *node) ([]*node, bool) {
	var inserted bool

	prefix := rune(word[0])
	suffix := word[1:]

	if len(nodes) == 0 {
		node, _ := newNode(prefix, suffix, parent)
		nodes = append(nodes, node)
		inserted = true
	}

	return nodes, inserted
}
