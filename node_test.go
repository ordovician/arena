package arenatree

import (
	"testing"
)

func TestCreateTreeNode(t *testing.T) {
	left := TreeNode[string, int]{Key: "four", Value: 4}
	right := TreeNode[string, int]{Key: "eight", Value: 8}

	root := TreeNode[string, int]{
		Key:   "forthytwo",
		Value: 42,
		left:  &left,
		right: &right,
	}

	if root.left.Value != 4 || root.right.Value != 8 {
		t.Errorf("Failed to create tree node")
	}
}
