package day07

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var cardStrengths2 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
	"J": 0,
}

func solvePart2(input string) int {
	// Implement your solution for part 2
	var sum int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cards := NewCards2(lines)
	sortedCards := sortCard2(cards)
	for rank, card := range sortedCards {
		rank = rank + 1
		sum += rank * card.Bid
	}

	return sum
}

func sortCard2(cards []*Card) []*Card {
	sort.Slice(cards, func(i, j int) bool {
		current := cards[i]
		next := cards[j]

		if current.TypeValue > next.TypeValue {
			return false
		}

		if current.TypeValue < next.TypeValue {
			return true
		}

		// if same values -> check the card strength
		for i, c := range current.id {
			currentValue := cardStrengths2[string(c)]
			nextValue := cardStrengths2[string(next.id[i])]
			if currentValue == nextValue {
				continue
			}

			return currentValue < nextValue
		}
		return false
	})

	return cards
}

func NewCards2(lines []string) []*Card {
	cards := make([]*Card, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		card := parts[0]
		bid := parts[1]
		intBid, err := strconv.Atoi(bid)
		if err != nil {
			log.Fatal("bid is not an int")
		}

		c := &Card{
			Bid: intBid,
			id:  card,
		}

		b := strings.Builder{}
		var countJokers int
		for c := range cardStrengths2 {
			v := strings.Count(card, c)
			if c == "J" {
				countJokers += v
				continue
			}

			// if we found one of a kind we don't care for it
			if v <= 1 {
				continue
			}

			b.WriteString(fmt.Sprintf("%d", v))
		}

		firstBuild := b.String()
		switch {
		case countJokers == 5 && firstBuild == "":
			b.WriteString(fmt.Sprintf("%d", countJokers))
		case countJokers > 0 && firstBuild != "":
			newString := getString(firstBuild, countJokers)
			b.Reset()
			b.WriteString(newString)
		case countJokers > 0 && firstBuild == "":
			b.WriteString(fmt.Sprintf("%d", countJokers+1))
		}

		c.Type = b.String()
		c.TypeValue = typeStrings[c.Type]
		cards = append(cards, c)
	}

	return cards
}

func getString(in string, countJokers int) string {
	var out string
	r := string(in[0])
	d, err := strconv.Atoi(r)
	if err != nil {
		log.Fatal("r is not an int")
	}
	d += countJokers
	if len(in) > 1 {
		n := string(in[1])
		out = fmt.Sprintf("%d%s", d, n)
	} else {
		out = fmt.Sprintf("%d", d)
	}
	return out
}
