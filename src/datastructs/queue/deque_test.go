package queue

import "testing"

func TestDeque(t *testing.T) {
	a := "A"
	b := "B"
	c := "C"

	q := NewDeque()
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

func BenchmarkDeque(b *testing.B) {
	q := NewDeque()

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
