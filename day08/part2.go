package day08

import (
	"strings"
)

func solvePart2(input string) int {
	// Implement your solution for part 2
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := lines[0]
	parsed := NewMap(lines[2:])
	startingNodes := findStartingPositions(parsed)

	return getSteps2(instructions, parsed, startingNodes)
}

func getSteps2(instructions string, parsed map[string][]string, startingNodes []string) int {
	steps := make([]int, 0)
	for _, node := range startingNodes {
		_, step := recursive2(node, instructions, 0, parsed)
		steps = append(steps, step)
	}

	return LCM(steps[0], steps[1], steps[1:]...)
}

func recursive2(node, instructions string, steps int, lookup map[string][]string) (string, int) {
	byteDirection := instructions[steps]
	direction := directionMap[byteDirection]

	nextNode := lookup[node][direction]
	steps++

	// found the end, break recursion
	if strings.HasSuffix(nextNode, "Z") {
		return nextNode, steps
	}

	// if we reach the end of the instructions,
	// double them and start over.
	if steps == len(instructions) {
		return recursive2(nextNode, instructions+instructions, steps, lookup)
	}

	return recursive2(nextNode, instructions, steps, lookup)
}

// LCM https://en.wikipedia.org/wiki/Least_common_multiple#Using_the_greatest_common_divisor
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// GCD https://en.wikipedia.org/wiki/Euclidean_algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func findStartingPositions(in map[string][]string) []string {
	startingPositions := make([]string, 0)

	for k := range in {
		if strings.HasSuffix(k, "A") {
			startingPositions = append(startingPositions, k)
		}
	}
	return startingPositions
}
