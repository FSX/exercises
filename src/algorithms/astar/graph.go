package astar

// Node is a multi-functional node!
type Node struct {
	Id string
	// X, Y int
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

// type SquareGrid struct {
// 	Width, Height int
// 	Walls         map[*Node]interface{}
// }

// func NewSquareGrid(width, height int) *SquareGrid {
// 	return &SquareGrid{
// 		width,
// 		height,
// 		make(map[int]interface{})}
// }

// func (s *SquareGrid) InBounds(n *Node) bool {
// 	return 0 <= n.X < s.Width && 0 <= n.Y < s.Height
// }

// func (s *SquareGrid) Passable(n *Node) bool {
// 	_, ok := s.Walls[n.X*n.Y]
// 	return !ok
// }

// // def neighbors(self, id):
// //     (x, y) = id
// //     results = [(x+1, y), (x, y-1), (x-1, y), (x, y+1)]
// //     if (x + y) % 2 == 0: results.reverse() # aesthetics
// //     results = filter(self.in_bounds, results)
// //     results = filter(self.passable, results)
// //     return results

// func (s *SquareGrid) Neighbors(n *Node) []*Node {
// 	x, y := n.X, n.Y

// 	// Usage of 'c' function is just for readability.
// 	r := []int{c(x+1, y), c(x, y-1), c(x-1, y), c(x, y+1)}

// 	if (x+y)%2 == 0 {
// 		r[0], r[1], r[2], r[3] = r[3], r[2], r[1], r[0]
// 	}

// 	// var results []int
// }

// func c(x, y int) int {
// 	return x * y
// }
