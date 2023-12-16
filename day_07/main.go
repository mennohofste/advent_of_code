package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var isPart2 = false

func readLines() []string {
	content, err := os.ReadFile("day_7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content = content[:len(content)-1] // remove trailing newline
	return strings.Split(string(content), "\n")
}

type Card rune

func (c Card) Value() int {
	if isPart2 && c == 'J' {
		return 1
	}

	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(c - '0')
	}
}

func (c Card) String() string {
	return string(c)
}

type Hand struct {
	Cards []Card
	Bid   int
}

func getHands(lines []string) []Hand {
	var hands []Hand
	for _, line := range lines {
		fields := strings.Fields(line)
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		hands = append(hands, Hand{Cards: []Card(fields[0]), Bid: bid})
	}
	return hands
}

func sortHands(hands []Hand) []Hand {
	slices.SortFunc(hands, func(a, b Hand) int {
		if n := cmp.Compare(a.Type(), b.Type()); n != 0 {
			return -n
		}
		for i := range a.Cards {
			if n := cmp.Compare(a.Cards[i].Value(), b.Cards[i].Value()); n != 0 {
				return n
			}
		}
		return 0
	})
	return hands
}

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func (h Hand) Type() HandType {
	cardCounts := map[Card]int{}
	for _, card := range h.Cards {
		cardCounts[card]++
	}

	mostCommonCard := Card('J')
	mostCommonCount := 0
	for card, count := range cardCounts {
		if count > mostCommonCount && card != 'J' {
			mostCommonCard = card
			mostCommonCount = count
		}
	}

	if isPart2 && mostCommonCount > 0 {
		cardCounts[mostCommonCard] += cardCounts['J']
		delete(cardCounts, 'J')
	}

	var handType HandType
	switch {
	case isFiveOfAKind(cardCounts):
		handType = FiveOfAKind
	case isFourOfAKind(cardCounts):
		handType = FourOfAKind
	case isFullHouse(cardCounts):
		handType = FullHouse
	case isThreeOfAKind(cardCounts):
		handType = ThreeOfAKind
	case isTwoPair(cardCounts):
		handType = TwoPair
	case isOnePair(cardCounts):
		handType = OnePair
	default:
		handType = HighCard
	}

	return handType
}

func isFiveOfAKind(cardCounts map[Card]int) bool {
	for _, count := range cardCounts {
		if count >= 5 {
			return true
		}
	}
	return false
}

func isFourOfAKind(cardCounts map[Card]int) bool {
	for _, count := range cardCounts {
		if count >= 4 {
			return true
		}
	}
	return false
}

func isFullHouse(cardCounts map[Card]int) bool {
	hasThree := false
	hasTwo := false
	for _, count := range cardCounts {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func isThreeOfAKind(cardCounts map[Card]int) bool {
	for _, count := range cardCounts {
		if count >= 3 {
			return true
		}
	}
	return false
}

func isTwoPair(cardCounts map[Card]int) bool {
	pairCount := 0
	for _, count := range cardCounts {
		if count >= 2 {
			pairCount++
		}
	}
	return pairCount >= 2
}

func isOnePair(cardCounts map[Card]int) bool {
	for _, count := range cardCounts {
		if count >= 2 {
			return true
		}
	}
	return false
}

func part1() int {
	lines := readLines()
	hands := getHands(lines)
	sortedHands := sortHands(hands)

	totalWinnnings := 0
	for rank, hand := range sortedHands {
		totalWinnnings += (rank + 1) * hand.Bid
	}

	return totalWinnnings
}

func part2() int {
	isPart2 = true
	return part1()
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
