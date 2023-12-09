package day08

import (
	"strings"
)

var directionMap = map[byte]int{
	byte(82): 1, // right
	byte(76): 0, // left
}

func solvePart1(input string) int {
	// Implement your solution for part 1
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := lines[0]
	parsed := NewMap(lines[2:])

	return getSteps(parsed, instructions)
}

func getSteps(in map[string][]string, instructions string) int {
	var position int
	startingNode := "AAA"

	_, steps := recursive(startingNode, instructions, position, in)
	return steps
}

func recursive(node, instructions string, steps int, lookup map[string][]string) (string, int) {
	byteDirection := instructions[steps]
	direction := directionMap[byteDirection]

	nextNode := lookup[node][direction]
	steps++

	// found the end, break recursion
	if nextNode == "ZZZ" {
		return nextNode, steps
	}

	// if we reach the end of the instructions,
	// double them and start over.
	if steps == len(instructions) {
		return recursive(nextNode, instructions+instructions, steps, lookup)
	}

	return recursive(nextNode, instructions, steps, lookup)
}

func NewMap(lines []string) map[string][]string {
	m := make(map[string][]string)

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " = ")
		left := parts[0]
		m[left] = parseNode(parts[1])
	}

	return m
}

func parseNode(node string) []string {
	out := make([]string, 0)
	tmp := strings.Split(node, ", ")
	out = append(out, tmp[0][1:])
	out = append(out, tmp[1][:len(tmp[1])-1])
	return out
}
