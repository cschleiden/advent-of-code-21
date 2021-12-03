package main

import "testing"

func Test_Driver(t *testing.T) {
	d := NewDriver([]string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	})

	s := NewSubmarine()
	d.Drive(s)

	if s.Depth() != 10 {
		t.Fail()
	}

	if s.Pos() != 15 {
		t.Fail()
	}
}
