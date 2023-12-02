package day02

import (
	"log"
	"strconv"
	"strings"
)

func solvePart1(input string) int {
	// Implement your solution for part 1
	config := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	games := make([]string, len(input))

	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		tmpSplit := strings.Split(line, ":")
		game := strings.Split(tmpSplit[0], " ")[1]

		found := false
		sets := strings.Split(tmpSplit[1], ";")

		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				tmpSplit := strings.Split(strings.TrimSpace(cube), " ")

				countStr, color := tmpSplit[0], tmpSplit[1]
				count := convertStrToInt(countStr)

				if v, ok := config[color]; ok {
					if count > v {
						found = true
						break
					}
				}
			}
		}

		if found {
			games[i] = ""
			continue
		}
		games[i] = game
	}

	var sum int
	for _, g := range games {
		if g == "" {
			continue
		}
		sum += convertStrToInt(g)
	}

	return sum
}

func convertStrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
