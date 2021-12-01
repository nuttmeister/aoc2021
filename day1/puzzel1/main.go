package main

import (
	"fmt"

	"github.com/nuttmeister/aoc2021/day1/internal/input"
)

func main() {
	increases := 0

	for i := range input.Depths {
		if i > 0 {
			if input.Depths[i] > input.Depths[i-1] {
				increases++
			}
		}
	}

	fmt.Printf("number of increases: %d\n", increases)
}
