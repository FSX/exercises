package queue

const (
	BLOCKLEN = 64
	CENTER   = (BLOCKLEN - 1) / 2
)

type qblock struct {
	next, prev *qblock
	data       [BLOCKLEN]interface{}
}

// Deque is based on collections.deque [1] from Python.
//
// A list of functions. On the left Deque functions,
// and on the right collections.deque functions:
//
// - Put = deque_append
// - Get = deque_popleft
//
// [1]: https://hg.python.org/cpython/file/tip/Modules/_collectionsmodule.c
type Deque struct {
	leftblock, rightblock        *qblock
	leftindex, rightindex, count int
}

func NewDeque() *Deque {
	b := &qblock{}
	q := &Deque{
		leftblock:  b,
		rightblock: b,
		leftindex:  CENTER + 1,
		rightindex: CENTER,
	}

	return q
}

// Put puts a value on the right side of the queue.
func (q *Deque) Put(v interface{}) {
	if q.rightindex == BLOCKLEN-1 {
		b := &qblock{}
		b.prev = q.rightblock
		q.rightblock.next = b
		q.rightblock = b
		q.rightindex = -1
	}

	q.rightindex++
	q.rightblock.data[q.rightindex] = v
	q.count++
}

// Get gets a value from the left side of the queue or nil if empty.
func (q *Deque) Get() interface{} {
	if q.count == 0 {
		return nil
	}

	v := q.leftblock.data[q.leftindex]
	q.leftindex++
	q.count--

	if q.leftindex == BLOCKLEN {
		if q.count > 0 {
			prev := q.leftblock.next
			q.leftblock = nil
			q.leftblock = prev
			q.leftindex = 0
		} else {
			q.leftindex = CENTER + 1
			q.rightindex = CENTER
		}
	}

	return v
}
