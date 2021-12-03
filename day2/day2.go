package main

import "fmt"

func main() {
	i := ReadInput()

	d := NewDriver(i)

	s := NewSubmarine()

	d.Drive(s)

	fmt.Println(s.Depth() * s.Pos())
}
