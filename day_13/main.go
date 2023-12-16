package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getLines() []string {
	bytes, err := os.ReadFile("day_13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.Trim(string(bytes), "\n"), "\n")
}

func getPatterns(lines []string) [][]string {
	var patterns [][]string
	var pattern []string
	for _, line := range lines {
		if line != "" {
			pattern = append(pattern, line)
		} else {
			patterns = append(patterns, pattern)
			pattern = []string{}
		}
	}
	patterns = append(patterns, pattern)
	return patterns
}

func distance(a, b string) int {
	if len(a) != len(b) {
		log.Fatal("Different lengths")
	}

	sum := 0
	for i := range a {
		if a[i] != b[i] {
			sum++
		}
	}

	return sum
}

func isMirrorIndex(field []string, index int, smudges int) bool {
	sum := 0
	for i, j := index-1, index; i >= 0 && j < len(field); i, j = i-1, j+1 {
		sum += distance(field[i], field[j])
		// optional optimization, prune early
		if sum > smudges {
			return false
		}
	}
	return sum == smudges
}

func mirrorIndex(field []string, smudges int) int {
	for i := 1; i < len(field); i++ {
		if isMirrorIndex(field, i, smudges) {
			return i
		}
	}
	return 0
}

func transpose(field []string) []string {
	transposed := make([]string, len(field[0]))
	for i := range transposed {
		for _, line := range field {
			transposed[i] += string(line[i])
		}
	}
	return transposed
}

func getSummary(patterns [][]string, smudges int) int {
	sum := 0
	for _, pattern := range patterns {
		horizontalValue := mirrorIndex(pattern, smudges)
		sum += horizontalValue * 100

		transposedLines := transpose(pattern)
		verticalValue := mirrorIndex(transposedLines, smudges)
		sum += verticalValue
	}
	return sum
}

func part1() int {
	lines := getLines()
	patterns := getPatterns(lines)
	return getSummary(patterns, 0)
}

func part2() int {
	lines := getLines()
	patterns := getPatterns(lines)
	return getSummary(patterns, 1)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
