package main

import (
	"fmt"

	"github.com/nuttmeister/aoc2021/day2/internal/input"
)

func main() {
	hor, depth := 0, 0

	for _, cmd := range input.Commands {
		switch cmd.Direction {
		case input.Up:
			depth = depth - cmd.Steps
		case input.Down:
			depth = depth + cmd.Steps
		case input.Forward:
			hor = hor + cmd.Steps
		}
	}

	fmt.Printf("horizontal: %d depth: %d result: %d\n", hor, depth, hor*depth)
}
