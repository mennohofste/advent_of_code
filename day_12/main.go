package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLines() []string {
	bytes, err := os.ReadFile("day_12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.Trim(string(bytes), "\n"), "\n")
}

type SpringLayout struct {
	Status string
	Config []int
}

func getSpringLayout(line string) SpringLayout {
	status, end, found := strings.Cut(line, " ")
	if !found {
		log.Fatal("invalid input")
	}

	strNumbers := strings.Split(end, ",")
	var config []int
	for _, strNumber := range strNumbers {
		number, err := strconv.Atoi(strNumber)
		if err != nil {
			log.Fatal(err)
		}
		config = append(config, number)
	}

	return SpringLayout{Status: status, Config: config}
}

func totalSize(blocks []int) int {
	sum := 0
	for _, size := range blocks {
		sum += size
	}
	return sum + len(blocks) - 1
}

var cache = make(map[string]int)

func getCache(symbols string, blocks []int) (int, bool) {
	var builder strings.Builder
	builder.WriteString(symbols)
	for _, block := range blocks {
		builder.WriteString(",")
		builder.WriteString(strconv.Itoa(block))
	}

	v, ok := cache[builder.String()]
	return v, ok
}

func setCache(symbols string, blocks []int, value int) {
	var builder strings.Builder
	builder.WriteString(symbols)
	for _, block := range blocks {
		builder.WriteString(",")
		builder.WriteString(strconv.Itoa(block))
	}

	cache[builder.String()] = value
}

func countConfigs(symbols string, blocks []int) int {
	if v, ok := getCache(symbols, blocks); ok {
		return v
	}

	// Termination conditions
	switch {
	case len(blocks) == 0:
		if strings.ContainsRune(symbols, '#') {
			return 0
		}
		return 1
	case len(symbols) < totalSize(blocks):
		return 0
	}

	switch symbols[0] {
	case '.':
		number := countConfigs(symbols[1:], blocks)
		setCache(symbols, blocks, number)
		return number
	case '#':
		for i := 0; i < blocks[0]; i++ {
			if symbols[i] == '.' || len(symbols) > blocks[0] && symbols[blocks[0]] == '#' {
				return 0
			}
		}
		if len(symbols) == blocks[0] && len(blocks) == 1 {
			return 1
		}
		number := countConfigs(symbols[blocks[0]+1:], blocks[1:])
		setCache(symbols, blocks, number)
		return number
	case '?':
		poundNumber := countConfigs(strings.Replace(symbols, "?", "#", 1), blocks)
		dotNumber := countConfigs(strings.Replace(symbols, "?", ".", 1), blocks)
		setCache(symbols, blocks, poundNumber+dotNumber)
		return poundNumber + dotNumber
	}

	log.Fatal("invalid input")
	return 0
}

func part1() int {
	lines := getLines()

	sum := 0
	for _, line := range lines {
		springLayout := getSpringLayout(line)
		sum += countConfigs(springLayout.Status, springLayout.Config)
	}

	return sum
}

func part2() int {
	lines := getLines()

	sum := 0
	for _, line := range lines {
		springLayout := getSpringLayout(line)
		newStatus := springLayout.Status
		newConfig := springLayout.Config
		for i := 0; i < 4; i++ {
			newStatus += "?" + springLayout.Status
			newConfig = append(newConfig, springLayout.Config...)
		}
		sum += countConfigs(newStatus, newConfig)
	}

	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
