package astar

// Node is a multi-functional node!
type Node struct {
	Id string
	// X, Y int
}

type Graph struct {
	edges map[string][]*Node
}

func NewGrap(edges map[string][]*Node) *Graph {
	return &Graph{edges}
}

func (g *Graph) Neighbors(n *Node) []*Node {
	if v, ok := g.edges[n.Id]; ok {
		return v
	} else {
		return []*Node{}
	}
}
