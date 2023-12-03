package day03

import (
	"fmt"
	"strings"
	"unicode"
)

func solvePart2(input string) int {
	// Implement your solution for part 2
	var sum int

	lines := strings.Split(strings.TrimSpace(input), "\n")
	coordinates := mapCoordinates(lines)
	// numbers := make(map[string]int, 0)
	numbers := make([]int, 0)

	for y, line := range lines {
		if line == "" {
			continue
		}

		for x, char := range line {
			// if the current char is not a star, skip
			if string(char) != "*" {
				continue
			}

			toCheck := []string{
				createCoordinates(x-1, y),   // left
				createCoordinates(x+1, y),   // right
				createCoordinates(x, y-1),   // up
				createCoordinates(x, y+1),   // down
				createCoordinates(x-1, y-1), // left up
				createCoordinates(x+1, y+1), // right down
				createCoordinates(x-1, y+1), // left down
				createCoordinates(x+1, y-1), // right up
			}

			prevChecks := make(map[string]int, 0)
			for _, c := range toCheck {
				v, ok := coordinates[c]
				// if the current char does exist, e.g. edges etc -> skip
				if !ok {
					continue
				}
				// if the coordinate doesn't have an adjacent number -> skip
				if !unicode.IsDigit(v) {
					continue
				}

				// found an adjacent number, now go back to that line
				// and find the full number by iterating over it again.
				tmpSplit := strings.Split(c, ",")
				xx, yy := strToDigit(tmpSplit[0]), strToDigit(tmpSplit[1])

				line := lines[yy]
				b := strings.Builder{}

				// find the start of the number
				var start int
				for i := xx; i >= 0; i-- {
					rr := rune(line[i])
					if !unicode.IsDigit(rr) {
						break
					}
					start = i
				}

				// find the full number
				for i := start; i < len(line); i++ {
					rr := rune(line[i])
					if !unicode.IsDigit(rr) {
						break
					}
					b.WriteRune(rr)
				}

				number := b.String()
				if number == "" {
					continue
				}

				prev := fmt.Sprintf("%s,%d", number, yy)
				// if we find the value in the same y again -> skip
				if _, ok := prevChecks[prev]; ok {
					continue
				}

				prevChecks[prev]++
			}
			// if we don't have 2 adjacent numbers -> skip
			if len(prevChecks) != 2 {
				continue
			}

			ratio := 1
			for k := range prevChecks {
				tmpSplit := strings.Split(k, ",")
				gear := strToDigit(tmpSplit[0])
				ratio *= gear
			}
			numbers = append(numbers, ratio)
		}
	}

	for _, v := range numbers {
		sum += v
	}

	return sum
}
