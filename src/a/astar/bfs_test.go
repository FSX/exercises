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
	// t.Log(neighbors)
	result := TracePath(neighbors, goal)
	t.Log(joinNodeIds(result, " -> "))
	// expected := []*Node{a, b, c, d, f}

	// if !reflect.DeepEqual(result, expected) {
	// 	t.Errorf(
	// 		"{%s} != {%s}",
	// 		joinNodeIds(result, ", "),
	// 		joinNodeIds(expected, ", "))
	// }

	t.Logf("\n%s", graph)
}
