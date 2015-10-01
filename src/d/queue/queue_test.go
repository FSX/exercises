package queue

import "testing"

func TestQueue(t *testing.T) {
	a := "A"
	b := "B"
	c := "C"

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

func BenchmarkQueue(b *testing.B) {
	q := &Queue{}

	for i := 0; i < 1000000; i++ {
		n := string(i)
		q.Put(n)
	}

	for i := 0; i < 1000000; i++ {
		if r := q.Get(); r == nil {
			b.Errorf("%v is nil, expected *node\n", r)
		}
	}
}
