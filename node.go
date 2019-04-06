package trie

// Node represents a node that contains a value and pointers to the parent and children
type Node interface {
	Value() rune
	Parent() Node
}

type node struct {
	value rune
}

// NewNode initializes a node with the value
func NewNode(value rune) Node {
	return &node{value}
}

func (n *node) Value() rune {
	return n.value
}

func (n *node) Parent() Node {
	return nil
}
