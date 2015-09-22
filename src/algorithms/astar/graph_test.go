package astar

import "testing"

func TestGraph(t *testing.T) {
	n := &Node{"A"}
	g := NewGrap()
	g.Neighbors(n)
}
