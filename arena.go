package arenatree

// Arena allocator. Useful for allocating lots of small objects of the same size
type Arena[T any] struct {
	blocks Stack[*T]
}

// Allocate a block of memory b which must be returned with arena.Free(b)
// when no longer used
func (arena *Arena[T]) Alloc() *T {
	if arena.blocks.IsEmpty() {
		// allocate in chunks of 8 to avoid too frequent allocations
		var blocks [8]T
		// No, you cannot iterate over blocks, because that would make a copy of the block
		// and you would get the address of the copy rather than address of blocks[i]
		// hence index i must be used when iterating
		for i, _ := range blocks {
			arena.blocks.Push(&blocks[i])
		}
	}
	b, _ := arena.blocks.Top()
	arena.blocks.Pop()

	return b
}

// Free block of memory previously allocated by arena allocator
func (arena *Arena[T]) Free(block *T) {
	if block == nil {
		panic("Cannot free nil pointer")
	}

	// A very basic test to see if we are doing a double-free
	// does not protect against freeing object further down the stack
	if top, ok := arena.blocks.Top(); !ok || top == block {
		panic("Releasing previosly freed pointer")
	}
	arena.blocks.Push(block)
}
