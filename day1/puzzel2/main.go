package main

import (
	"fmt"

	"github.com/nuttmeister/aoc2021/day1/internal/input"
)

func main() {
	increases := 0

	for i := range input.Depths {
		if i > 2 {
			prev := input.Depths[i-3] + input.Depths[i-2] + input.Depths[i-1]
			cur := input.Depths[i-2] + input.Depths[i-1] + input.Depths[i]
			if cur > prev {
				increases++
			}
		}
	}

	fmt.Printf("number of increases: %d\n", increases)
}
