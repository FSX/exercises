package observer

type Event int

const (
	EventOne Event = iota
	EventTwo
	EventThree
)

type Observer interface {
	Receive(entity interface{}, event Event)
}

type Subject struct {
	head *node
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) notify(entity interface{}, event Event) {
	for c := s.head; c != nil; c = c.n {
		c.v.Receive(entity, event)
	}
}

func (s *Subject) Add(o Observer) {
	n := &node{v: o}
	s.head, n.n = n, s.head
}

func (s *Subject) Remove(o Observer) bool {
	if s.head.v == o {
		s.head, s.head.n = s.head.n, nil
		return true
	}

	for c := s.head; c != nil; c = c.n {
		if n := c.n; n != nil && n.v == o {
			c.n, n = n.n, nil
			return true
		}
	}

	return false
}

type node struct {
	v Observer
	n *node
}
