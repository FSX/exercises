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

	return reversePath(path)
}

func reversePath(path []*Node) []*Node {
	a := make([]*Node, len(path))
	l := len(path) - 1
	m := (l + 1) / 2

	if l%2 == 0 {
		a[m] = path[m]
	}
	for i := 0; i < m; i++ {
		a[i], a[l-i] = path[l-i], path[i]
	}

	return a
}
