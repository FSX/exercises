package main

import (
	"d/queue"
	"github.com/pkg/profile"
)

func fill(q *queue.Deque) {
	s := 1024 * 1024 * 10

	for i := 0; i < 34; i++ {
		b := make([]byte, s)
		for n := 0; n < s; n++ {
			b[n] = byte(n)
		}

		q.Put(b)
	}
}

func clear(q *queue.Deque) {
	for i := 0; i < 34; i++ {
		v := q.Get()
		if v == nil {
			panic("value is nil\n")
		}
		v = nil
	}
}

func main() {
	p := profile.Start(
		profile.MemProfile,
		profile.ProfilePath("."))
	defer p.Stop()

	q := queue.NewDeque()
	fill(q)
	clear(q)
}
