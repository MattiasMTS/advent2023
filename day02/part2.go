package day02

import (
	"strings"
)

func solvePart2(input string) int {
	// Implement your solution for part 2
	games := make([]map[string]int, len(input))

	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		tmpSplit := strings.Split(line, ":")

		sets := strings.Split(tmpSplit[1], ";")
		config := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				tmpSplit := strings.Split(strings.TrimSpace(cube), " ")

				countStr, color := tmpSplit[0], tmpSplit[1]
				count := convertStrToInt(countStr)

				if v, ok := config[color]; ok {
					if count > v {
						config[color] = count
					}
				}
			}
		}
		games[i] = config
	}

	var sum int
	for _, game := range games {
		// ignore the empty games from making the slice
		if len(game) == 0 {
			continue
		}

		power := 1
		for _, v := range game {
			power *= v
		}
		sum += power
	}

	return sum
}
