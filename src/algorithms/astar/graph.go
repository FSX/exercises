package astar

type Graph interface {
	Neighbors(n *Node) []*Node
}

type Node struct {
	Id string
	// X, Y int
}

type DirectedGraph struct {
	edges map[string][]*Node
}

func NewDirectedGraph(edges map[string][]*Node) *DirectedGraph {
	return &DirectedGraph{edges}
}

func (d *DirectedGraph) Neighbors(n *Node) []*Node {
	if v, ok := d.edges[n.Id]; ok {
		return v
	} else {
		return []*Node{}
	}
}
