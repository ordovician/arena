package arenatree

import "golang.org/x/exp/constraints"

type Tree[K constraints.Ordered, V any] struct {
	Root      *TreeNode[K, V]
	allocator Arena[TreeNode[K, V]]
}

func NewTree[K constraints.Ordered, V any]() *Tree[K, V] {
	var tree Tree[K, V]
	return &tree
}

func (tree *Tree[K, V]) NewNode(key K, value V) *TreeNode[K, V] {
	n := tree.allocator.Alloc()
	n.Key = key
	n.Value = value
	n.left = nil
	n.right = nil

	return n
}

func (tree *Tree[K, V]) Find(key K) (V, bool) {
	var (
		notFound V
		ok       bool
		node     **TreeNode[K, V]
	)
	node, ok = tree.Root.findNode(key)
	if !ok {
		return notFound, false
	}
	return (*node).Value, true
}

func (tree *Tree[K, V]) release(n *TreeNode[K, V]) {
	if n.left != nil {
		tree.release(n.left)
	}
	if n.right != nil {
		tree.release(n.right)
	}
	n.left = nil
	n.right = nil
	tree.allocator.Free(n)
}

func (tree *Tree[K, V]) Insert(key K, value V) {
	n := tree.NewNode(key, value)

	if tree.Root == nil {
		tree.Root = n
	} else {
		tree.Root.Insert(n)
	}
}

func (tree *Tree[K, V]) Delete(key K) {
	node, ok := tree.Root.findNode(key)
	if !ok {
		return
	}

	tree.release(*node)
	*node = nil
}

func (tree *Tree[K, V]) Iterator() <-chan *TreeNode[K, V] {
	return tree.Root.Iterator()
}
