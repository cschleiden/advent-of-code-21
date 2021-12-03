package main

type Submarine interface {
	Up(v int)
	Down(v int)
	Forward(v int)
	Backward(v int)

	Depth() int
	Pos() int
}

func NewSubmarine() Submarine {
	return &submarine{}
}

type submarine struct {
	depth, pos, aim int
}

func (s *submarine) Up(v int) {
	s.aim -= v
}

func (s *submarine) Down(v int) {
	s.aim += v
}

func (s *submarine) Forward(v int) {
	s.pos += v
	s.depth += v * s.aim
}

func (s *submarine) Backward(v int) {
	s.pos -= v
}

func (s *submarine) Depth() int {
	return s.depth
}

func (s *submarine) Pos() int {
	return s.pos
}
