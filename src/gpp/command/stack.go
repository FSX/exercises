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
	size       int
	head, tail *commandNode
}

func NewCommandStack(size int) *CommandStack {
	node := &commandNode{command: &noopCommand{}}
	return &CommandStack{size, node, node}
}

func (c *CommandStack) Do(command Command) {
	node := &commandNode{command: command}

	c.tail.next = node
	node.prev = c.tail
	c.tail = node

	// Limit list of actions to c.size.
	i := 1
	n := c.tail

	for n != nil {
		if _, ok := n.command.(*noopCommand); ok {
			break
		} else if i == c.size {
			n.prev = c.head
			n.next = nil
			break
		}

		n = n.prev
		i++
	}

	command.Execute()
}

func (c *CommandStack) Undo() error {
	if c.tail.prev == nil {
		return ErrUndo
	}

	c.tail.command.Undo()
	c.tail = c.tail.prev

	return nil
}

func (c *CommandStack) Redo() error {
	if c.tail.next == nil {
		return ErrRedo
	}

	c.tail = c.tail.next
	c.tail.command.Execute()

	return nil
}

func (c *CommandStack) Len() int {
	i := 0

	for n := c.tail; n != nil; n = n.prev {
		i++
	}
	for n := c.tail.next; n != nil; n = n.next {
		i++
	}

	return i - 1 // noopCommand shouldn't be included in the count.
}

type commandNode struct {
	command    Command
	next, prev *commandNode
}
