package main

import (
	"io/ioutil"
	"strings"
)

func ReadInput() []string {
	i, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic("could not read input")
	}

	return strings.Split(string(i), "\n")
}
