package main

import (
	"strconv"
	"strings"
)

type driver struct {
	commands []string
}

type Driver interface {
	Drive(s Submarine)
}

func NewDriver(commands []string) Driver {
	return &driver{
		commands: commands,
	}
}

func (d *driver) Drive(s Submarine) {
	for _, c := range d.commands {
		p := strings.Split(c, " ")

		v, _ := strconv.Atoi(p[1])

		switch p[0] {
		case "up":
			s.Up(v)
		case "down":
			s.Down(v)
		case "forward":
			s.Forward(v)
		case "Backward":
			s.Backward(v)
		}
	}
}
