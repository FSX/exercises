package command

import "errors"

var (
	ErrUndo = errors.New("no undo commands left")
	ErrRedo = errors.New("no redo commands left")
)

type noopCommand struct{}

func (n *noopCommand) Execute()       {}
func (n *noopCommand) Undo()          {}
func (n *noopCommand) String() string { return "noop()" }

type CommandStack struct {
	size           int
	first, current *commandNode
}

func NewCommandStack(size int) *CommandStack {
	node := &commandNode{command: &noopCommand{}}
	return &CommandStack{size, node, node}
}

func (c *CommandStack) Do(command Command) {
	node := &commandNode{command: command}

	c.current.next = node
	node.prev = c.current
	c.current = node

	// Limit list of actions to c.size.
	i := 1
	n := c.current

	for n != nil {
		if _, ok := n.command.(*noopCommand); ok {
			break
		} else if i == c.size {
			n.prev = c.first
			break
		}

		n = n.prev
		i++
	}

	command.Execute()
}

func (c *CommandStack) Undo() error {
	if c.current.prev == nil {
		return ErrUndo
	}

	c.current.command.Undo()
	c.current = c.current.prev

	return nil
}

func (c *CommandStack) Redo() error {
	if c.current.next == nil {
		return ErrRedo
	}

	c.current = c.current.next
	c.current.command.Execute()

	return nil
}

func (c *CommandStack) Len() int {
	i := 0

	for n := c.current; n != nil; n = n.prev {
		i++
	}
	for n := c.current.next; n != nil; n = n.next {
		i++
	}

	return i - 1 // noopCommand shouldn't be included in the count.
}

type commandNode struct {
	command    Command
	next, prev *commandNode
}
