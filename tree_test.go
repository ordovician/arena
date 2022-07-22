package arenatree

import (
	"testing"
)

func createTestTree() Tree[int, string] {
	var tree Tree[int, string]
	tree.Insert(8, "eight")
	tree.Insert(4, "four")
	tree.Insert(10, "ten")

	return tree
}

func TestTreeAllocation(t *testing.T) {
	var tree Tree[int, string]

	var nodes []*TreeNode[int, string] = []*TreeNode[int, string]{
		tree.NewNode(1, "one"),
		tree.NewNode(4, "four"),
		tree.NewNode(2, "two"),
		tree.NewNode(9, "nine"),
	}

	for i, n := range nodes {
		for j, m := range nodes {
			if i == j {
				continue
			}
			if n == m {
				t.Errorf("Node with key %v got allocated twice!", n.Key)
			}
		}
	}

}

func TestCreateTreeNodeFromArena(t *testing.T) {
	var tree Tree[int, string]

	tree.Insert(8, "eight")

	if tree.Root.Key != 8 {
		t.Errorf("Got %v, but expected 8", tree.Root.Key)
	}

	tree.Insert(4, "four")

	if tree.Root.left.Key != 4 {
		t.Errorf("Got %v, but expected 4", tree.Root.left.Key)
	}

	tree.Insert(10, "ten")

	if tree.Root.right.Key != 10 {
		t.Errorf("Got %v, but expected 10", tree.Root.right.Key)
	}

}

func TestFindTreeNode(t *testing.T) {
	tree := createTestTree()

	eight, _ := tree.Find(8)
	four, _ := tree.Find(4)
	ten, _ := tree.Find(10)

	if eight != "eight" {
		t.Errorf("Got %v wanted eight", eight)
	}

	if four != "four" {
		t.Errorf("Got %v wanted four", four)
	}

	if ten != "ten" {
		t.Errorf("Got %v wanted ten", ten)
	}
}

func TestDeleteTreeNode(t *testing.T) {
	tree := createTestTree()

	eight, _ := tree.Root.findNode(8)

	if node, found := tree.allocator.blocks.Top(); found && node == *eight {
		t.Errorf("Node with key %v has been allocated and should not be in free list", (*eight).Key)
	}

	tree.Delete(8)
	if node, found := tree.allocator.blocks.Top(); found && node != *eight {
		t.Errorf("Just released node %v should be at top of free list", (*eight).Key)
	}

}
