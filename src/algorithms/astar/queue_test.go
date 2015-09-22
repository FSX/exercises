package astar

import "testing"

func TestQueue(t *testing.T) {
	a := &Node{"A"}
	b := &Node{"B"}
	c := &Node{"C"}

	q := &Queue{}
	q.Put(a)
	q.Put(b)
	q.Put(c)

	if r := q.Get(); r != a {
		t.Errorf("%v != %v\n", r, a)
	}
	if r := q.Get(); r != b {
		t.Errorf("%v != %v\n", r, b)
	}
	if r := q.Get(); r != c {
		t.Errorf("%v != %v\n", r, c)
	}
	if r := q.Get(); r != nil {
		t.Errorf("%v != nil\n", r)
	}
}
