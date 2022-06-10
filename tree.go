package arenatree

import "golang.org/x/exp/constraints"

type Tree[K constraints.Ordered, V any] struct {
	Root *TreeNode[K, V]
	free *TreeNode[K, V] // list of released nodes ready to be allocated
}

func (tree *Tree[K, V]) NewNode(key K, value V) *TreeNode[K, V] {
	if tree.free == nil {
		nodes := make([]TreeNode[K, V], 4)
		for i := 0; i < len(nodes)-1; i++ {
			nodes[i].left = &nodes[i+1]
		}
		tree.free = &nodes[0]
	}
	n := tree.free
	tree.free = tree.free.left

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
	n.left = tree.free
	n.right = nil
	tree.free = n
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
