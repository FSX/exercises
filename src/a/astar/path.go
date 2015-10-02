package astar

func TracePath(nodeMap map[*Node]*Node, goal *Node) []*Node {
	path := []*Node{goal}

	for {
		n, ok := nodeMap[goal]
		if ok && n != nil {
			goal = n
			path = append(path, n)
		} else {
			break
		}
	}

	return reverseNodes(path)
}

func reverseNodes(path []*Node) []*Node {
	newPath := make([]*Node, len(path))

	for a, b := 0, len(path)-1; a < len(path); a, b = a+1, b-1 {
		newPath[a] = path[b]
	}

	return newPath
}
