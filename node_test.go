package trie

import "testing"

func TestNewNode(t *testing.T) {
	node, inserted := newNode('f', "un", nil)

	if !inserted {
		t.Error("inserted should be true")
	}

	if value := node.value; value != 'f' {
		t.Errorf("value must be 'f'; found %#U", value)
	}

	if node.parent != nil {
		t.Error("the node must not have a parent")
	}

	if node.endOfWord {
		t.Error("this node is not the end of word")
	}

	if len(node.children) != 1 {
		t.Error("the node must have one child")
	}

	if node.children[0].value != 'u' {
		t.Errorf("value must be 'u'; found %#U", node.children[0].value)
	}

	if node.children[0].parent != node {
		t.Error("The parent of the child must be the same node")
	}

	if node.children[0].endOfWord {
		t.Error("this node is not the end of word")
	}

	if len(node.children[0].children) != 1 {
		t.Error("the node must have one child")
	}

	if node.children[0].children[0].value != 'n' {
		t.Errorf("value must be 'n'; found %#U", node.children[0].value)
	}

	if node.children[0].children[0].parent != node.children[0] {
		t.Error("The parent of the child must be the same node")
	}

	if !node.children[0].children[0].endOfWord {
		t.Error("this node is the end of word")
	}
}
