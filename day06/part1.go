package day06

import (
	"strconv"
	"strings"
)

type Document struct {
	time     int
	distance int
}

func solvePart1(input string) int {
	// Implement your solution for part 1
	lines := strings.Split(strings.TrimSpace(input), "\n")

	games := getGames1(lines)
	return calculateWinningWays(games)
}

func calculateWinningWays(games map[int]*Document) int {
	out := 1

	for _, doc := range games {
		var time int
		for i := 0; i <= doc.time; i++ {
			timeToMove := doc.time - i
			endDistance := timeToMove * i
			if endDistance > doc.distance {
				time++
			}
		}
		out *= time
	}

	return out
}

func getGames1(lines []string) map[int]*Document {
	games := make(map[int]*Document)
	for _, line := range lines {
		var game int

		if line == "" {
			continue
		}

		if strings.Contains(line, "Time:") {
			tmp := strings.Split(line, " ")
			for _, n := range tmp {
				d, err := strconv.Atoi(n)
				if err != nil {
					continue
				}
				games[game] = &Document{time: d}
				game++
			}
		}

		if strings.Contains(line, "Distance:") {
			tmp := strings.Split(line, " ")
			for _, n := range tmp {
				d, err := strconv.Atoi(n)
				if err != nil {
					continue
				}
				games[game].distance = d
				game++
			}
		}
	}

	return games
}
