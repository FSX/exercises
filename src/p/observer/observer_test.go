package observer

import "testing"

type TestObserver struct {
	Name      string
	Counter   int
	LastEvent Event
}

func (t *TestObserver) Receive(entity interface{}, event Event) {
	t.LastEvent = event
	t.Counter++
}

func TestAdd(t *testing.T) {
	o1 := &TestObserver{Name: "1"}
	o2 := &TestObserver{Name: "2"}

	s := NewSubject()
	s.Add(o1)
	s.Add(o2)

	i := 0
	for c := s.head; c != nil; c = c.n {
		i++
	}

	if i != 2 {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	o1 := &TestObserver{Name: "1"}
	o2 := &TestObserver{Name: "2"}
	o3 := &TestObserver{Name: "3"}
	o4 := &TestObserver{Name: "4"}

	s := NewSubject()
	s.Add(o1)
	s.Add(o2)
	s.Add(o3)

	if !s.Remove(o3) {
		t.Fail()
	}

	if s.Remove(o4) {
		t.Fail()
	}

	i := 0
	for c := s.head; c != nil; c = c.n {
		i++
	}

	if i != 2 {
		t.Fail()
	}
}

func TestNotify(t *testing.T) {
	o1 := &TestObserver{Name: "1"}
	s := NewSubject()
	s.Add(o1)

	s.notify(s, EventOne)

	if o1.Counter != 1 || o1.LastEvent != EventOne {
		t.Fail()
	}
}
