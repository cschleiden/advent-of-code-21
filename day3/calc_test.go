package main

import (
	"fmt"
	"testing"
)

func Test_SmallInput(t *testing.T) {
	gamma, epsilon := calc([]string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	})

	if gamma != 22 || epsilon != 9 {
		fmt.Println(gamma, epsilon)
		t.Fail()
	}
}

func Test_SmallInput2(t *testing.T) {
	oxygen, scrubber := calc_2([]string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	})

	if oxygen != 23 || scrubber != 10 {
		fmt.Println(oxygen, scrubber)
		t.Fail()
	}
}
