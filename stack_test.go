package arenatree

import (
	"testing"

	// to get convenient assert function
	"github.com/stretchr/testify/assert"
)

func TestCreatingStack(t *testing.T) {
	var stack Stack[int]

	stack.Push(42)

	if x, ok := stack.Top(); !ok || x != 42 {
		t.Errorf("Got %v from top of stack, expected 42", x)
	}
	assert.Equal(t, 1, stack.Len(), "Not expected number of elements in stack")

	stack.Push(1331)
	if x, ok := stack.Top(); !ok || x != 1331 {
		t.Errorf("Got %v from top of stack, expected 1331", x)
	}
	assert.Equal(t, 2, stack.Len(), "Not expected number of elements in stack")
}

func TestPoppingStack(t *testing.T) {
	var stack Stack[int]

	stack.Push(2)
	stack.Push(4)
	stack.Push(6)

	assert.Equal(t, 3, stack.Len(), "Not expected number of elements")

	stack.Pop()
	stack.Pop()
	stack.Pop()

	assert.Equal(t, 0, stack.Len(), "Not expected number of elements")
}

func TestPoppingEmptyStack(t *testing.T) {
	var stack Stack[int]

	defer func() {
		if r := recover(); r == nil {
			t.Error("Popping empty stack should cause panic")
		}
	}()

	stack.Pop()
}
