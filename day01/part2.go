package day01

import (
	"strings"
)

func solvePart2(input string) int {
	// Implement your solution for part 2
	var out int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		out += parseLine(line)
	}

	return out
}
