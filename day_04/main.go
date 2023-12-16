package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Scratchcard struct {
	WinNumbers []int
	OurNumbers []int
	Id         int
}

func (s Scratchcard) Score() int {
	score := 0
	for _, winningNumber := range s.WinNumbers {
		for _, ourNumber := range s.OurNumbers {
			if winningNumber == ourNumber {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}
	return score
}

func (s Scratchcard) Wins() int {
	wins := 0
	for _, winningNumber := range s.WinNumbers {
		for _, ourNumber := range s.OurNumbers {
			if winningNumber == ourNumber {
				wins++
			}
		}
	}
	return wins
}

func getScratchcards() []Scratchcard {
	content, err := os.ReadFile("day_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Remove last newline
	content = content[:len(content)-1]
	lines := strings.Split(string(content), "\n")

	var scratchcards []Scratchcard
	for _, line := range lines {
		items := strings.Fields(line)

		// Second item without the last colon
		cardId, err := strconv.Atoi(items[1][:len(items[1])-1])
		if err != nil {
			log.Fatal(err)
		}

		var winNumbers []int
		var ourNumbers []int
		isWin := true
		for _, item := range items[2:] {
			if item == "|" {
				isWin = false
				continue
			}

			number, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			if isWin {
				winNumbers = append(winNumbers, number)
			} else {
				ourNumbers = append(ourNumbers, number)
			}
		}

		scratchcards = append(scratchcards, Scratchcard{Id: cardId, WinNumbers: winNumbers, OurNumbers: ourNumbers})
	}
	return scratchcards
}

func part1() int {
	scratchcards := getScratchcards()
	sum := 0
	for _, scratchcard := range scratchcards {
		sum += scratchcard.Score()
	}
	return sum
}

func part2() int {
	scratchcards := getScratchcards()

	scratchCopies := make([]int, len(scratchcards))
	for i := range scratchcards {
		scratchCopies[i] = 1
	}

	for _, scratchcard := range scratchcards {
		for i := 1; i <= scratchcard.Wins(); i++ {
			scratchCopies[scratchcard.Id+i-1] += scratchCopies[scratchcard.Id-1]
		}
	}

	sum := 0
	for _, copies := range scratchCopies {
		sum += copies
	}
	return sum
}

func main() {
	println("Part 1:", part1())
	println("Part 2:", part2())
}
