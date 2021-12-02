package main

import "fmt"

func increases(input []int) int {
	increases := 0

	previous := -1
	for _, depth := range input1 {
		if previous != -1 && depth > previous {
			increases = increases + 1
		}

		previous = depth
	}

	return increases
}

func sliding_window_increases(input []int) int {
	var increases int
	previous := -1

	for start := range input {
		if start+2 >= len(input) {
			break
		}

		w := input[start] + input[start+1] + input[start+2]

		if previous != -1 && previous < w {
			increases = increases + 1
		}

		previous = w
	}

	return increases
}

func main() {
	fmt.Println("Increases", increases(input1))
	fmt.Println("Increases Sliding Window", sliding_window_increases(input1))
}
