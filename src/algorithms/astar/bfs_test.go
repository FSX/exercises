package astar

import (
	"reflect"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	a := &Node{"A"}
	b := &Node{"B"}
	c := &Node{"C"}
	d := &Node{"D"}
	e := &Node{"E"}
	f := &Node{"F"}
	g := &Node{"G"}

	graph := NewDirectedGraph(map[*Node][]*Node{
		a: {b},
		b: {c},
		c: {d},
		d: {e, f},
		e: {c},
		f: {g},
		g: {a},
	})

	goal := f
	neighbors := BreadthFirstSearch(graph, a, goal)
	result := TracePath(neighbors, goal)
	expected := []*Node{a, b, c, d, f}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"{%s} != {%s}",
			joinNodeIds(result, ", "),
			joinNodeIds(expected, ", "))
	}
}
