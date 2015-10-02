package astar

type Graph interface {
	Neighbors(*Node) []*Node
}

type Node struct {
	Id   string
	X, Y int
}

type DirectedGraph struct {
	edges map[*Node][]*Node
}

func NewDirectedGraph(edges map[*Node][]*Node) *DirectedGraph {
	return &DirectedGraph{edges}
}

func (d *DirectedGraph) Neighbors(node *Node) []*Node {
	if v, ok := d.edges[node]; ok {
		return v
	} else {
		return []*Node{}
	}
}
