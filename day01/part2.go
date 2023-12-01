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
		digits := make([]string, 0)

		for i := 0; i < len(line); i++ {
			if line[i] >= '1' && line[i] <= '9' {
				digits = append(digits, string(line[i]))
				continue
			}

			for k, v := range digitMap {
				if i+len(k) > len(line) {
					continue
				}
				if line[i:i+len(k)] != k {
					continue
				}
				digits = append(digits, v)
			}
		}
		comb := combineDigits(digits)
		out += convertToDigit(comb)
	}

	return out
}
