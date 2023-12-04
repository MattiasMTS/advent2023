package day04

import (
	"log"
	"strings"
)

func solvePart2(input string) int {
	// Implement your solution for part 2
	var sum int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cards := processLines(lines)

	for _, card := range cards {
		sum += card
	}

	return sum
}

func processLines(lines []string) []int {
	out := make([]int, len(lines))
	for i, line := range lines {
		if line == "" {
			log.Fatal("line is empty")
		}

		countWinners := getNumberWinners(line)
		count := countWinners + i

		// if we are at the end of the lines, we can't extend further
		if count >= len(lines) {
			count = len(lines) - 1
		}

		// initialize the card
		out[i] += 1

		// extend with copies of the cards found
		for j := i + 1; j <= count; j++ {
			out[j] += out[i]
		}
	}

	return out
}

func getCard(line string) string {
	return strings.TrimSpace(strings.Split(line, ":")[0])
}

func getNumberWinners(line string) int {
	cardSplit := strings.Split(line, ":")

	numbers := strings.Split(cardSplit[len(cardSplit)-1], "|")
	if len(numbers) != 2 {
		log.Fatal("splitting on input and winner cards failed")
	}

	var foundCards int
	winners := strings.Split(numbers[0], " ")
	input := strings.Split(numbers[1], " ")

	for _, card := range input {
		if card == "" {
			continue
		}

		if !IsIn(card, winners) {
			continue
		}

		foundCards++
	}

	return foundCards
}
