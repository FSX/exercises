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
	max           int
	length        int
	head, current *commandNode
}

func NewCommandStack(max int) *CommandStack {
	node := &commandNode{command: &noopCommand{}}
	return &CommandStack{max, 0, node, node}
}

func (c *CommandStack) Do(command Command) {
	node := &commandNode{command: command}

	// Append node at the end.
	c.current.next = node
	node.prev = c.current
	c.current = node

	// Count nodes.
	newLength := 1
	for n := c.current; n != nil; n = n.prev {
		newLength++
	}

	// Truncate everything in front of the new node.
	if newLength > c.max {
		c.head.next = c.head.next.next
		c.head.next.prev = c.head
		c.length = newLength - 1
	} else {
		c.length = newLength
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
	return c.length
	// i := 0

	// for n := c.current; n != nil; n = n.prev {
	// 	i++
	// }
	// for n := c.current.next; n != nil; n = n.next {
	// 	i++
	// }

	// return i - 1 // noopCommand shouldn't be included in the count.
}

type commandNode struct {
	command    Command
	next, prev *commandNode
}
