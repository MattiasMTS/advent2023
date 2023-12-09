package day09

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func solvePart1(input string) int {
	// Implement your solution for part 1
	var sum int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		sum += recursive1(line, nil)
	}

	return sum
}

func recursive1(line string, lastValues []int) int {
	var (
		sum     int
		diff    int
		current int
		next    int
	)
	values := strings.Split(line, " ")
	b := strings.Builder{}

	for i, v := range values {
		if v == "" {
			continue
		}

		if i+1 == len(values) {
			break
		}

		current, next = getDigits1(v, values, i)
		diff = next - current
		sum += diff

		b.WriteString(fmt.Sprintf(" %d", diff))
	}
	lastValues = append(lastValues, next)

	// if sum is 0 we have found the diff and break the loop
	if sum == 0 {
		for _, v := range lastValues {
			sum += v
		}
		return sum
	}

	return recursive1(b.String(), lastValues)
}

func getDigits1(v string, values []string, i int) (int, int) {
	current, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}
	next, err := strconv.Atoi(values[i+1])
	if err != nil {
		log.Fatal(err)
	}
	return current, next
}
