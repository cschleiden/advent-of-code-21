package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("could not read input")
	}

	gamma, epsilon := calc(strings.Split(string(inputs), "\n"))
	fmt.Println(gamma, epsilon, gamma*epsilon)

	oxygen, scrubber := calc_2(strings.Split(string(inputs), "\n"))
	fmt.Println(oxygen, scrubber, oxygen*scrubber)
}
