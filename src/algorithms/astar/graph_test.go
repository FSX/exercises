package astar

import (
	"reflect"
	"testing"
)

type neighborTest struct {
	Root      *Node
	Neighbors []*Node
}

func TestGraph(t *testing.T) {
	a := &Node{"A"}
	b := &Node{"B"}
	c := &Node{"C"}
	d := &Node{"D"}
	e := &Node{"E"}
	f := &Node{"F"}
	g := &Node{"G"}

	graph := NewGrap(map[string][]*Node{
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
	if len(nodes) == 0 {
		return ""
	}
	if len(nodes) == 1 {
		return nodes[0].Id
	}

	n := len(sep) * (len(nodes) - 1)
	for i := 0; i < len(nodes); i++ {
		n += len(nodes[i].Id)
	}

	b := make([]byte, n)
	bp := copy(b, nodes[0].Id)

	for _, node := range nodes[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], node.Id)
	}

	return string(b)
}
