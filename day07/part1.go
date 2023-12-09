package day07

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var cardStrengths1 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var typeStrings = map[string]int{
	"":   0, // all cards are the same
	"2":  1, // one pair
	"22": 2, // two pair
	"3":  3, // three of a kind
	"32": 4, // full house
	"23": 4, // full house
	"4":  5, // four of a kind
	"5":  6, // five of a kind
}

type Card struct {
	id        string
	Type      string
	TypeValue int
	Bid       int
}

func solvePart1(input string) int {
	// Implement your solution for part 1
	var sum int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cards := NewCards1(lines)
	sortedCards := sortCard1(cards)
	for rank, card := range sortedCards {
		rank = rank + 1
		sum += rank * card.Bid
	}

	return sum
}

func sortCard1(cards []*Card) []*Card {
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
			currentValue := cardStrengths1[string(c)]
			nextValue := cardStrengths1[string(next.id[i])]
			if currentValue == nextValue {
				continue
			}

			return currentValue < nextValue
		}
		return false
	})

	return cards
}

func NewCards1(lines []string) []*Card {
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
		for c := range cardStrengths1 {
			v := strings.Count(card, c)
			// if we found one of a kind we don't care for it
			if v <= 1 {
				continue
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}

		c.Type = b.String()
		c.TypeValue = typeStrings[c.Type]
		cards = append(cards, c)
	}

	return cards
}
