package arenatree

import "fmt"

func ExampleTree_Iterator() {
	tree := createTestTree()

	// This checks we are getting all the values
	channel := tree.Iterator()
	for i := 0; i < 3; i++ {
		node := <-channel
		fmt.Println(node.Key)
	}

	// checks that we terminate channel with a close
	for node := range tree.Iterator() {
		fmt.Println(node.Key)
	}

	// Output:
	// 8
	// 4
	// 10
	// 8
	// 4
	// 10
}
