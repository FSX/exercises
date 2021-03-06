package astar

import (
	"reflect"
	"strings"
	"testing"
)

func joinNodeIds(nodes []*Node, sep string) string {
	a := make([]string, len(nodes))

	for i := 0; i < len(nodes); i++ {
		a[i] = nodes[i].Id
	}

	return strings.Join(a, sep)
}

type neighborTest struct {
	Root      *Node
	Neighbors []*Node
}

func TestDirectedGraph(t *testing.T) {
	a := &Node{Id: "A"}
	b := &Node{Id: "B"}
	c := &Node{Id: "C"}
	d := &Node{Id: "D"}
	e := &Node{Id: "E"}
	f := &Node{Id: "F"}
	g := &Node{Id: "G"}

	graph := NewDirectedGraph(map[*Node][]*Node{
		a: {b, c},
		b: {d},
		c: {f, b, g},
		d: {e, f},
		e: {c},
		f: {g},
		g: {a},
	})

	tests := []neighborTest{
		{a, []*Node{b, c}},
		{b, []*Node{d}},
		{c, []*Node{f, b, g}},
		{d, []*Node{e, f}},
		{e, []*Node{c}},
		{f, []*Node{g}},
		{g, []*Node{a}},
	}

	for _, s := range tests {
		result := graph.Neighbors(s.Root)
		if !reflect.DeepEqual(result, s.Neighbors) {
			t.Errorf(
				"{%s} != {%s}",
				joinNodeIds(result, ", "),
				joinNodeIds(s.Neighbors, ", "))
		}
	}
}
