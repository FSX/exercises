package astar

type Node struct {
	Id string
}

type Graph struct {
	edges map[string][]*Node
}

func NewGrap() *Graph {
	return &Graph{edges: make(map[string][]*Node)}
}

func (g *Graph) Neighbors(n *Node) []*Node {
	if v, ok := g.edges[n.Id]; ok {
		return v
	} else {
		return []*Node{}
	}
}
