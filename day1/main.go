package main

import "fmt"

func main() {
	increases := 0

	previous := -1
	for _, depth := range input1 {
		if previous != -1 && depth > previous {
			increases = increases + 1
		}

		previous = depth
	}

	fmt.Println("Increases", increases)
}
