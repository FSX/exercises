package bitvector

import (
	// "fmt"
	"testing"
)

func TestIntSet(t *testing.T) {
	s := &IntSet{}
	s.Add(1)
	s.Add(144)
	s.Add(9)

	if !s.Has(1) {
		t.Fatal("set does not have 1")
	}
	if !s.Has(144) {
		t.Fatal("set does not have 144")
	}
	if !s.Has(9) {
		t.Fatal("set does not have 9")
	}
	if s.Has(999) {
		t.Fatal("set has 999")
	}
}

func TestIntSetAddAll(t *testing.T) {
	s := &IntSet{}
	s.AddAll(1, 144, 9)

	if !s.Has(1) {
		t.Fatal("set does not have 1")
	}
	if !s.Has(144) {
		t.Fatal("set does not have 144")
	}
	if !s.Has(9) {
		t.Fatal("set does not have 9")
	}
	if s.Has(999) {
		t.Fatal("set has 999")
	}
}

func TestIntSetUnion(t *testing.T) {
	a := &IntSet{}
	a.Add(55)
	b := &IntSet{}
	b.Add(66)

	a.UnionWith(b)

	if !a.Has(55) {
		t.Fatal("set does not have 55")
	}
	if !a.Has(66) {
		t.Fatal("set does not have 66")
	}
}

func TestIntSetLen(t *testing.T) {
	s := &IntSet{}
	s.Add(55)
	s.Add(56)
	s.Add(57)
	s.Add(58)
	s.Add(59)
	s.Add(60)

	if s.Len() != 6 {
		t.Fatal("length of set is not 6")
	}
}

func TestIntSetRemove(t *testing.T) {
	s := &IntSet{}
	s.Add(55)
	s.Add(56)
	s.Remove(55)

	if s.Has(55) {
		t.Fatal("set must not include 55")
	}
	if !s.Has(56) {
		t.Fatal("set must include 56")
	}
}

func TestIntSetClear(t *testing.T) {
	s := &IntSet{}
	s.Add(55)
	s.Add(56)

	if s.Len() != 2 {
		t.Fatal("set must have 2 elements")
	}

	s.Clear()

	if s.Len() != 0 {
		t.Fatal("set must have no elements")
	}

	s.Add(55)
	s.Add(56)

	if s.Len() != 2 {
		t.Fatal("set must have 2 elements")
	}
}

func TestIntSetCopy(t *testing.T) {
	s := &IntSet{}
	s.Add(55)
	s.Add(56)

	p := s.Copy()

	if !p.Has(55) {
		t.Fatal("set must include 55")
	}
	if !p.Has(56) {
		t.Fatal("set must include 56")
	}
}

func TestIntSetIntersectWith(t *testing.T) {
	a := &IntSet{}
	a.Add(55)
	a.Add(56)
	a.Add(99)
	b := &IntSet{}
	b.Add(55)
	b.Add(56)
	b.Add(100)

	a.IntersectWith(b)

	if !a.Has(55) {
		t.Fatal("set must include 55")
	}
	if !a.Has(56) {
		t.Fatal("set must include 56")
	}
	if a.Has(99) {
		t.Fatal("set must not include 99")
	}
	if a.Has(100) {
		t.Fatal("set must not include 100")
	}
}

func TestIntSetDifferenceWith(t *testing.T) {
	a := &IntSet{}
	a.Add(55)
	a.Add(56)
	a.Add(99)
	b := &IntSet{}
	b.Add(55)
	b.Add(56)
	b.Add(100)

	a.DifferenceWith(b)

	if a.Has(55) {
		t.Fatal("set must noy include 55")
	}
	if a.Has(56) {
		t.Fatal("set must noy include 56")
	}
	if !a.Has(99) {
		t.Fatal("set must include 99")
	}
	if a.Has(100) {
		t.Fatal("set must not include 100")
	}
}
