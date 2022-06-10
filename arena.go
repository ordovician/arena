package arenatree

// Arena allocator. Useful for allocating lots of small objects of the same size
type Arena[T any] struct {
	blocks Stack[*T]
}

// Allocate a block of memory b which must be returned with arena.Free(b)
// when no longer used
func (arena *Arena[T]) Alloc() *T {
	if arena.blocks.IsEmpty() {
		var blocks [4]T
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
	arena.blocks.Push(block)
}
