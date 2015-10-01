package astar

type Graph interface {
	Neighbors(n *Node) []*Node
}

type Node struct {
	Id string
	// X, Y int
}

type DirectedGraph struct {
	edges map[*Node][]*Node
}

func NewDirectedGraph(edges map[*Node][]*Node) *DirectedGraph {
	return &DirectedGraph{edges}
}

func (d *DirectedGraph) Neighbors(n *Node) []*Node {
	if v, ok := d.edges[n]; ok {
		return v
	} else {
		return []*Node{}
	}
}
