package day04

import (
	"log"
	"strings"
)

func solvePart1(input string) int {
	// Implement your solution for part 1
	var sum int

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		carSplit := strings.Split(line, ":")
		numbers := strings.Split(carSplit[len(carSplit)-1], "|")
		if len(numbers) != 2 {
			log.Fatal("splitting on input and winner cards failed")
		}

		winners := strings.Split(numbers[0], " ")
		input := strings.Split(numbers[1], " ")
		var rowCount int
		for _, card := range input {
			if card == "" {
				continue
			}

			if !IsIn(card, winners) {
				continue
			}

			if rowCount == 0 {
				rowCount++
			} else {
				rowCount *= 2
			}
		}

		sum += rowCount
	}
	return sum
}

func IsIn(input string, winners []string) bool {
	for _, winner := range winners {
		if winner == input {
			return true
		}
	}

	return false
}
