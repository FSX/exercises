package main

import (
	"d/queue"
	"github.com/pkg/profile"
)

func main() {
	p := profile.Start(
		profile.MemProfile,
		profile.ProfilePath("."))
	defer p.Stop()

	q := &queue.Queue{}

	for i := 0; i < 20000000; i++ {
		n := string(i)
		q.Put(n)
	}

	for i := 0; i < 20000000; i++ {
		if q.Get() == nil {
			panic("value is nil, expected *node\n")
		}
	}
}
