package astar

import (
	"reflect"
	"strings"
	"testing"
)

type neighborTest struct {
	Root      *Node
	Neighbors []*Node
}

func TestDirectedGraph(t *testing.T) {
	a := &Node{"A"}
	b := &Node{"B"}
	c := &Node{"C"}
	d := &Node{"D"}
	e := &Node{"E"}
	f := &Node{"F"}
	g := &Node{"G"}

	graph := NewDirectedGraph(map[string][]*Node{
		"A": {b, c},
		"B": {d},
		"C": {f, b, g},
		"D": {e, f},
		"E": {c},
		"F": {g},
		"G": {a},
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
				joinNodeIds(s.Neighbors, ", "),
				joinNodeIds(result, ", "))
		}
	}
}

func joinNodeIds(nodes []*Node, sep string) string {
	a := make([]string, len(nodes))

	for i := 0; i < len(nodes); i++ {
		a[i] = nodes[i].Id
	}

	return strings.Join(a, sep)
}
