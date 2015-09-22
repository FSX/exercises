package astar

type element struct {
	v          *Node // Graph nodes only.
	next, prev *element
}

type Queue struct {
	head *element
	tail *element
}

func (q *Queue) Empty() bool {
	return q.head == nil
}

func (q *Queue) Put(n *Node) {
	if q.head != nil {
		e := &element{v: n}
		q.tail.next = e
		e.prev = q.tail
		q.tail = e
	} else {
		q.head = &element{v: n}
		q.tail = q.head
		q.head.next = q.tail
		q.tail.prev = q.head
	}
}

func (q *Queue) Get() *Node {
	if q.head == nil {
		return nil
	}

	n := q.head.v
	q.head = q.head.next

	if q.head != nil {
		q.head.prev = nil
	}

	return n
}
