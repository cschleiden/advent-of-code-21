package main

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
