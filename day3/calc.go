package main

import "strconv"

func calc(inputs []string) (gamma, epsilon int) {
	width := len(inputs[0])

	c := make([]int, width)

	for _, bin := range inputs {
		for i := range c {
			if bin[i] == '1' {
				c[i]++
			}
		}
	}

	x := len(inputs) / 2

	for i, c := range c {
		if c >= x {
			gamma = (1 << (width - i - 1)) | gamma
		}
		if c <= x {
			epsilon = (1 << (width - i - 1)) | epsilon
		}
	}

	return gamma, epsilon
}

func calc_2(inputs []string) (oxygen, scrubber int64) {
	oxygen = calc_2_do(inputs, func(a int, b int) bool {
		return a >= b
	})

	scrubber = calc_2_do(inputs, func(a int, b int) bool {
		return a < b
	})

	return oxygen, scrubber
}

func calc_2_do(inputs []string, op func(a, b int) bool) int64 {
	candidates := inputs

	i := 0

	for len(candidates) > 1 {
		ones := make([]string, 0)
		zeroes := make([]string, 0)

		for _, bin := range candidates {
			if bin[i] == '1' {
				ones = append(ones, bin)
			} else {
				zeroes = append(zeroes, bin)
			}
		}

		if op(len(ones), len(zeroes)) {
			candidates = ones
		} else {
			candidates = zeroes
		}

		i++
	}

	v, err := strconv.ParseInt(candidates[0], 2, 32)
	if err != nil {
		panic("could not convert")
	}

	return v
}
