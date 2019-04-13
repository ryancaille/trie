package trie

import "sort"

type node struct {
	value     rune
	parent    *node
	children  []*node
	endOfWord bool
}

// create initializes a node with the value,
// and passes in the next suffix that will be inserted into its children
func create(r rune, suffix []rune, parent *node) (*node, bool) {

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

		nodeToInsert, _ := create(prefix, suffix, parent)

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
func contains(nodes []*node, word []rune) (bool, *node) {

	// Search for the node that matches the rune
	_, node := search(nodes, word[0])
	if node == nil {
		return false, nil
	}

	// if this is the last rune in the word, we don't search anymore children
	if endOfWord := len(word) == 1; endOfWord {
		if node.endOfWord {
			return true, node
		}
		return false, nil
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

func remove(rootChildren []*node, word []rune) ([]*node, bool) {

	found, n := contains(rootChildren, word)

	if found {

		n.endOfWord = false

		c := cleanup(n, n.parent)

		if c != nil {
			// delete the child from the parent's children
			return deleteChild(rootChildren, c), found
		}
	}

	return rootChildren, found
}

// cleanup remove a node from the parent if that node has no children
func cleanup(n *node, p *node) *node {

	// when p is nil, it means n is a root child and we need to remove it
	if p == nil {
		return n
	}

	// if n has children, cleanup is done
	if len(n.children) > 0 {
		return nil
	}

	// delete the child from the parent's children
	p.children = deleteChild(p.children, n)

	// if the parent is the end of another word, cleanup is done
	if p.endOfWord {
		return nil
	}

	// recurse up the tree to determine if the parent needs to be deleted
	return cleanup(p, p.parent)
}

func deleteChild(children []*node, child *node) []*node {
	if i, c := search(children, child.value); c == child {

		// remove the child c at index i, and preserve order
		copy(children[i:], children[i+1:])
		children[len(children)-1] = nil
		children = children[:len(children)-1]
	}

	return children
}
