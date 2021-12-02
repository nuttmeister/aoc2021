package main

import (
	"fmt"

	"github.com/nuttmeister/aoc2021/day2/internal/input"
)

func main() {
	aim, hor, depth := 0, 0, 0

	for _, cmd := range input.Commands {
		switch cmd.Direction {
		case input.Up:
			aim = aim - cmd.Steps
		case input.Down:
			aim = aim + cmd.Steps
		case input.Forward:
			hor = hor + cmd.Steps
			depth = depth + (aim * cmd.Steps)
		}
	}

	fmt.Printf("aim: %d horizontal: %d depth: %d result: %d\n", aim, hor, depth, hor*depth)
}
