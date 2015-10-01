package astar

import "d/queue"

func BreadthFirstSearch(graph Graph, start, goal *Node) map[*Node]*Node {
	frontier := queue.NewDeque()
	frontier.Put(start)
	cameFrom := make(map[*Node]*Node)
	cameFrom[start] = nil

	for !frontier.Empty() {
		if current, ok := frontier.Get().(*Node); ok {
			if current == goal {
				break
			}

			for _, next := range graph.Neighbors(current) {
				if _, ok := cameFrom[next]; !ok {
					frontier.Put(next)
					cameFrom[next] = current
				}
			}
		} else {
			panic("current is not *Node")
		}
	}

	return cameFrom
}
