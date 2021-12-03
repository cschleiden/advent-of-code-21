package main

import "testing"

func Test_BasicNavigation(t *testing.T) {
	s := NewSubmarine()

	s.Forward(5)
	s.Down(5)
	s.Forward(8)
	s.Up(3)
	s.Down(8)
	s.Forward(2)

	if s.Depth() != 60 {
		t.Fail()
	}

	if s.Pos() != 15 {
		t.Fail()
	}
}
