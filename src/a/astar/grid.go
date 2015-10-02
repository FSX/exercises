package astar

import (
	"bytes"
	"fmt"
	"strings"
)

type SquareGrid struct {
	Width, Height int
	nodes         [][]*Node
	walls         [][4]int
	dirs          [4][2]int
}

func NewSquareGrid(width, height int) *SquareGrid {
	nodes := make([][]*Node, height)

	for y := 0; y < height; y++ {
		nodes[y] = make([]*Node, width)

		for x := 0; x < width; x++ {
			nodes[y][x] = &Node{Id: fmt.Sprintf("(%d,%d)", x, y), X: x, Y: y}
		}
	}

	return &SquareGrid{
		Width:  width,
		Height: height,
		nodes:  nodes,
		dirs:   [4][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}},
	}
}

func (s *SquareGrid) GetNode(x, y int) (*Node, bool) {
	if !s.inBounds(x, y) || !s.passable(x, y) {
		return nil, false
	}

	return s.nodes[y][x], true
}

func (s *SquareGrid) AddWall(x, y, width, height int) {
	if !s.inBounds(x, y) || !s.inBounds(x+width, y+height) {
		panic("wall is not in bounds")
	}

	s.walls = append(s.walls, [4]int{x, y, x + width, y + height})
}

func (s *SquareGrid) inBounds(x, y int) bool {
	if x >= 0 && x < s.Width && y >= 0 && y < s.Height {
		return true
	}

	return false
}

func (s *SquareGrid) passable(x, y int) bool {
	for i := 0; i < len(s.walls); i++ {
		x1, y1, x2, y2 :=
			s.walls[i][0], s.walls[i][1],
			s.walls[i][2], s.walls[i][3]

		if x >= x1 && x < x2 && y >= y1 && y < y2 {
			return false
		}
	}

	return true
}

func (s *SquareGrid) Neighbors(node *Node) []*Node {
	x, y := node.X, node.Y
	result := make([]*Node, 0, 4)

	for i := 0; i < 4; i++ {
		nx, ny := x+s.dirs[i][0], y+s.dirs[i][1]

		if s.inBounds(nx, ny) && s.passable(nx, ny) {
			result = append(result, s.nodes[ny][nx])
		}
	}

	if (x+y)%2 == 0 {
		result = reverseNodes(result)
	}

	return result
}

func (s *SquareGrid) String() string {
	var b bytes.Buffer

	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			if s.passable(x, y) {
				n := s.nodes[y][x].Id

				b.WriteString(strings.Repeat(" ", 8-len(n)))
				b.WriteString(n)
			} else {
				b.WriteString("  ######")
			}
		}

		b.WriteRune('\n')
	}

	return b.String()
}
