package astar

type element struct {
	v          interface{}
	next, prev *element
}

type Queue struct {
	head *element
	tail *element
}

func (q *Queue) Empty() bool {
	return q.head == nil
}

func (q *Queue) Put(v interface{}) {
	if q.head != nil {
		e := &element{v: v}
		q.tail.next = e
		e.prev = q.tail
		q.tail = e
	} else {
		q.head = &element{v: v}
		q.tail = q.head
		q.head.next = q.tail
		q.tail.prev = q.head
	}
}

func (q *Queue) Get() interface{} {
	if q.head == nil {
		return nil
	}

	v := q.head.v
	q.head = q.head.next

	if q.head != nil {
		q.head.prev = nil
	}

	return v
}
