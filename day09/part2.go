package day09

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func solvePart2(input string) int {
	// Implement your solution for part 2
	var sum int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		sum += recursive2(line, nil)
	}

	return sum
}

func recursive2(line string, lastValues []int) int {
	var (
		sum     int
		diff    int
		current int
		next    int
	)
	values := strings.Split(line, " ")
	b := ""

	for i := len(values) - 1; i >= 0; i-- {
		v := values[i]
		if v == "" {
			continue
		}

		if i-1 < 0 {
			break
		}

		current, next = getDigits2(v, values, i)
		diff = current - next
		sum += diff

		b = fmt.Sprintf("%d ", diff) + b
	}
	// need to add the last value to the first position
	// since we are going in reverse
	lastValues = append(lastValues, next)

	// if sum is 0 we have found the diff and break the loop
	if sum == 0 {
		for i := len(lastValues) - 1; i >= 0; i-- {
			v := lastValues[i]
			sum = v - sum
		}
		return sum
	}

	return recursive2(b, lastValues)
}

func getDigits2(v string, values []string, i int) (int, int) {
	current, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}
	prev, err := strconv.Atoi(values[i-1])
	if err != nil {
		log.Fatal(err)
	}
	return current, prev
}
