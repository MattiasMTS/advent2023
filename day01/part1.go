package day01

import (
	"log"
	"strconv"
	"strings"
)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func solvePart1(input string) int {
	// Implement your solution for part 1
	var out int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		out += parseLine(line)
	}

	return out
}

func parseLine(line string) int {
	digits := make([]string, 0)

	for i := 0; i < len(line); i++ {
		if line[i] >= '1' && line[i] <= '9' {
			digits = append(digits, string(line[i]))
			continue
		}

		for k, v := range digitMap {
			if i+len(k) > len(line) {
				continue
			}
			if line[i:i+len(k)] != k {
				continue
			}
			digits = append(digits, v)
		}
	}
	comb := digits[0] + digits[len(digits)-1]
	if len(comb) > 2 {
		log.Fatal("digit too long")
	}
	return convertToDigit(comb)
}

func convertToDigit(in string) int {
	d, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
