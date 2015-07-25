package command

import (
	"fmt"
)

type Command interface {
	Execute()
	Undo()
}

type MoveUnitCommand struct {
	unit   *Unit
	cx, cy int // Current
	px, py int // Previous
}

func NewMoveUnitCommand(unit *Unit, x, y int) *MoveUnitCommand {
	return &MoveUnitCommand{unit, x, y, 0, 0}
}

func (m *MoveUnitCommand) Execute() {
	m.px, m.py = m.unit.Position()
	m.unit.MoveTo(m.cx, m.cy)
}

func (m *MoveUnitCommand) Undo() {
	m.unit.MoveTo(m.px, m.py)
}

func (m *MoveUnitCommand) String() string {
	return fmt.Sprintf("move(%d,%d)", m.cx, m.cy)
}
