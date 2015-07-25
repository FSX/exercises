package command

import (
	"fmt"
	"testing"
)

func TestUndo(*testing.T) {
	u := &Unit{"Test-01", 0, 0}
	s := NewCommandStack(10)

	s.Do(NewMoveUnitCommand(u, 10, 10))
	s.Do(NewMoveUnitCommand(u, 20, 20))
	s.Do(NewMoveUnitCommand(u, 30, 30))
	s.Do(NewMoveUnitCommand(u, 40, 40))
	s.Do(NewMoveUnitCommand(u, 50, 50))
	s.Do(NewMoveUnitCommand(u, 60, 60))
	s.Do(NewMoveUnitCommand(u, 70, 70))
	s.Do(NewMoveUnitCommand(u, 80, 80))
	s.Do(NewMoveUnitCommand(u, 90, 90))
	s.Do(NewMoveUnitCommand(u, 100, 100))
	s.Do(NewMoveUnitCommand(u, 110, 110))
	s.Do(NewMoveUnitCommand(u, 120, 120))

	for n := s.current; n != nil; n = n.prev {
		fmt.Println(n.command)
	}

	fmt.Println(s.Count())

}
