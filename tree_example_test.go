package arenatree

import "fmt"

func ExampleTreeTraversal() {

	tree := createTestTree()

	for node := range tree.Iterator() {
		fmt.Println(node.Key)
	}

	// Output:
}
