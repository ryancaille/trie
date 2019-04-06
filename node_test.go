package trie

import "testing"

func TestNewNode(t *testing.T) {
	node := NewNode('f')

	if value := node.Value(); value != 'f' {
		t.Errorf("value must be 'f'; found %#U", value)
	}

	if node.Parent() != nil {
		t.Error("the node must not have a parent")
	}
}
