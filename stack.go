package arenatree

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

func (stack *Stack[T]) Pop() {
	n := len(stack.items)
	if n <= 0 {
		panic("Cannot pop an empty stack!")
	}
	stack.items = stack.items[:n-1]
}

func (stack *Stack[T]) Top() (T, bool) {
	var top T
	n := len(stack.items)
	if n == 0 {
		return top, false
	}
	return stack.items[n-1], true
}

func (stack *Stack[T]) Len() int {
	return len(stack.items)
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.Len() == 0
}
