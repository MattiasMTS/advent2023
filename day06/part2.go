package day06

import (
	"log"
	"strconv"
	"strings"
)

func solvePart2(input string) int {
	// Implement your solution for part 1
	lines := strings.Split(strings.TrimSpace(input), "\n")

	games := getGames2(lines)
	return calculateWinningWays(games)
}

func getGames2(lines []string) map[int]*Document {
	games := make(map[int]*Document)
	for _, line := range lines {
		var game int

		if line == "" {
			continue
		}

		if strings.Contains(line, "Time:") {
			d := getNumber(line)
			games[game] = &Document{time: d}
			game++
		}

		if strings.Contains(line, "Distance:") {
			d := getNumber(line)
			games[game].distance = d
			game++
		}
	}

	return games
}

func getNumber(line string) int {
	tmp := strings.Split(line, " ")
	b := strings.Builder{}
	for _, n := range tmp {
		_, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		b.WriteString(n)
	}
	d, err := strconv.Atoi(b.String())
	if err != nil {
		log.Fatal(err)
	}
	return d
}
