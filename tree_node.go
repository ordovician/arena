package arenatree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[K constraints.Ordered, V any] struct {
	Key         K
	Value       V
	left, right *TreeNode[K, V]
}

func (n *TreeNode[K, V]) String() string {
	return fmt.Sprintf("TreeNode(%v, %v)", n.Key, n.Value)
}

// Insert node n into a leaf underneath parent node.
// Position will be determined based on value of key
func (parent *TreeNode[K, V]) Insert(n *TreeNode[K, V]) {
	if n.Key >= parent.Key {
		if parent.right == nil {
			parent.right = n
		} else {
			parent.right.Insert(n)
		}
	} else {
		if parent.left == nil {
			parent.left = n
		} else {
			parent.left.Insert(n)
		}
	}
}

func (n *TreeNode[K, V]) findNode(key K) (**TreeNode[K, V], bool) {
	if key == n.Key {
		return &n, true
	} else if key > n.Key {
		if n.right != nil {
			return n.right.findNode(key)
		}
	} else {
		if n.left != nil {
			return n.left.findNode(key)
		}
	}

	return nil, false
}

func (parent *TreeNode[K, V]) Iterator() <-chan *TreeNode[K, V] {
	channel := make(chan *TreeNode[K, V])

	go func() {
		parent.traverse(channel)
	}()

	return channel
}

func (n *TreeNode[K, V]) traverse(out chan<- *TreeNode[K, V]) {
	out <- n
	if n.left != nil {
		n.left.traverse(out)
	}
	if n.right != nil {
		n.right.traverse(out)
	}
}
