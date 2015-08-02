package command

import "testing"

func TestMoveUnitCommand(t *testing.T) {
	u := &Unit{"Test-01", 0, 0}

	c := NewMoveUnitCommand(u, 10, 10)
	c.Execute()

	// Check if unit moved.
	if x, y := u.Position(); x != 10 || y != 10 {
		t.Fail()
	}

	// Check if previous location is stored.
	c = NewMoveUnitCommand(u, 20, 20)
	c.Execute()

	if c.px != 10 || c.py != 10 {
		t.Fail()
	}
}

func TestStackLimit(t *testing.T) {
	u := &Unit{"Test-02", 0, 0}
	s := NewCommandStack(10)

	for i := 1; i <= 20; i++ {
		s.Do(NewMoveUnitCommand(u, i*10, i*10))
	}

	if l := s.Len(); l != 10 {
		t.Fail()
		t.Errorf("expected 10, but is %d", l)
	}
}

func TestUndo(t *testing.T) {
	u := &Unit{"Test-03", 0, 0}
	s := NewCommandStack(10)

	s.Do(NewMoveUnitCommand(u, 10, 10))
	s.Do(NewMoveUnitCommand(u, 20, 20))
	s.Do(NewMoveUnitCommand(u, 30, 30))
	s.Do(NewMoveUnitCommand(u, 40, 40))
	s.Do(NewMoveUnitCommand(u, 50, 50))

	if x, y := u.Position(); x != 50 || y != 50 {
		t.Fail()
	}

	// Two commands back.
	s.Undo()
	s.Undo()

	if x, y := u.Position(); x != 30 || y != 30 {
		t.Fail()
	}

	// Three commands back.
	s.Undo()
	s.Undo()
	s.Undo()

	if x, y := u.Position(); x != 0 || y != 0 {
		t.Fail()
	}

	// Another undo is not possible.
	if err := s.Undo(); err != ErrUndo {
		t.Fail()
	}
}

func TestRedo(t *testing.T) {
	u := &Unit{"Test-03", 0, 0}
	s := NewCommandStack(10)

	s.Do(NewMoveUnitCommand(u, 10, 10))
	s.Do(NewMoveUnitCommand(u, 20, 20))
	s.Do(NewMoveUnitCommand(u, 30, 30))
	s.Do(NewMoveUnitCommand(u, 40, 40))
	s.Do(NewMoveUnitCommand(u, 50, 50))

	s.Undo()
	s.Undo()
	s.Undo()

	if x, y := u.Position(); x != 20 || y != 20 {
		t.Fail()
	}

	s.Redo()
	s.Redo()
	s.Redo()

	if x, y := u.Position(); x != 50 || y != 50 {
		t.Fail()
	}

	// Another redo is not possible.
	if err := s.Redo(); err != ErrRedo {
		t.Fail()
	}
}

func TestNewAction(t *testing.T) {
	u := &Unit{"Test-02", 0, 0}
	s := NewCommandStack(10)

	for i := 1; i <= 10; i++ {
		s.Do(NewMoveUnitCommand(u, i*10, i*10))
	}

	// Go back 5 commands.
	for i := 1; i <= 5; i++ {
		s.Undo()
	}

	if s.Len() != 10 {
		t.Fail()
	}

	// Execute a new command. Every command that's newer than the current
	// command will be removed, before the following command is inserted.
	s.Do(NewMoveUnitCommand(u, 110, 110))

	// This means the new length is 6.
	if s.Len() != 6 {
		t.Fail()
	}
}
