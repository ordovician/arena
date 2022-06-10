package arenatree

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	// to get convenient assert function
)

type Point struct {
	x, y int
}

func TestArenaAllocation(t *testing.T) {
	var arena Arena[Point]

	p := arena.Alloc()
	q := arena.Alloc()

	p.x = 1
	p.y = 2
	q.x = 3
	q.y = 4

	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", q)

	assert.NotEqual(t, p, q, "p and q should point to different objects")
	assert.NotEqual(t, *p, *q, "p and q should have different content")

}

func TestArenaFree(t *testing.T) {
	var arena Arena[Point]

	p := arena.Alloc()
	arena.Free(p)
	q := arena.Alloc()
	r := arena.Alloc()

	// Use old fashion checks because asser.Equal did not work properly. So don't
	// put you faith in assert.Equal it maybe overpromising.
	if p != q {
		t.Errorf("p: %p should not be equal to q: %p since q is a reuse of p", p, q)
	}

	if p == r {
		t.Errorf("p: %p should not be equal to r: %p since r is not a reuse of p", p, r)
	}
}
