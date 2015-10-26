package astar

import (
	"reflect"
	"testing"
)

func TestBreadthFirstSearchDirectedGraph(t *testing.T) {
	a := &Node{Id: "A"}
	b := &Node{Id: "B"}
	c := &Node{Id: "C"}
	d := &Node{Id: "D"}
	e := &Node{Id: "E"}
	f := &Node{Id: "F"}
	g := &Node{Id: "G"}

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

func TestBreadthFirstSearchSquareGrid(t *testing.T) {
	graph := NewSquareGrid(20, 20)
	graph.AddWall(10, 5, 1, 5)
	graph.AddWall(15, 8, 1, 5)

	start, ok := graph.GetNode(10, 2)
	if !ok {
		t.Error("cannot get start node")
	}
	goal, ok := graph.GetNode(16, 12)
	if !ok {
		t.Error("cannot get goal node")
	}

	neighbors := BreadthFirstSearch(graph, start, goal)
	result := TracePath(neighbors, goal)

	// Make expected path.
	coords := [][]int{
		{10, 2}, {10, 3}, {11, 3}, {11, 4}, {12, 4}, {12, 5}, {13, 5},
		{13, 6}, {14, 6}, {14, 7}, {15, 7}, {16, 7}, {16, 8}, {16, 9},
		{16, 10}, {16, 11}, {16, 12},
	}
	expected := make([]*Node, len(coords))
	for i, c := range coords {
		if node, ok := graph.GetNode(c[0], c[1]); ok {
			expected[i] = node
		} else {
			t.Error("cannot get node")
		}
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"{%s} != {%s}",
			joinNodeIds(result, ", "),
			joinNodeIds(expected, ", "))
	}
}
