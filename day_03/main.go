package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Symbol struct {
	IsGear bool
	X      int
	Y      int
}

type Number struct {
	XStart int
	XEnd   int
	Y      int
	Value  int
}

func (n Number) NextTo(symbol Symbol) bool {
	return n.Y-symbol.Y <= 1 && symbol.Y-n.Y <= 1 && n.XStart-1 <= symbol.X && symbol.X <= n.XEnd+1
}

func getLines() []string {
	bytes, err := os.ReadFile("day_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.Trim(string(bytes), "\n"), "\n")
}

func getSymbols(schematic []string) []Symbol {
	var symbols []Symbol
	for i, line := range schematic {
		for j, char := range line {
			if char != '.' && !unicode.IsDigit(char) {
				symbols = append(symbols, Symbol{char == '*', j, i})
			}
		}
	}
	return symbols
}

func getNumbers(schematic []string) []Number {
	var numbers []Number
	for i, line := range schematic {
		numberStart := -1
		var numberChars []rune
		for j, char := range line {
			if unicode.IsDigit(char) {
				numberChars = append(numberChars, char)
				if numberStart == -1 {
					numberStart = j
				}
			} else if numberStart != -1 {
				numberValue, err := strconv.Atoi(string(numberChars))
				if err != nil {
					log.Fatal(err)
				}
				numbers = append(numbers, Number{numberStart, j - 1, i, numberValue})
				numberChars = []rune{}
				numberStart = -1
			}
		}

		if numberStart != -1 {
			numberValue, err := strconv.Atoi(string(numberChars))
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, Number{numberStart, len(line) - 1, i, numberValue})
		}
	}
	return numbers
}

func part1() int {
	schematic := getLines()
	symbols := getSymbols(schematic)
	numbers := getNumbers(schematic)

	sum := 0
	for _, number := range numbers {
		for _, symbol := range symbols {
			if number.NextTo(symbol) {
				sum += number.Value
				break
			}
		}
	}

	return sum
}

func part2() int {
	schematic := getLines()
	symbols := getSymbols(schematic)
	numbers := getNumbers(schematic)

	sum := 0
	for _, symbol := range symbols {
		if symbol.IsGear {
			var chosenNumbers []Number
			for _, number := range numbers {
				if number.NextTo(symbol) {
					chosenNumbers = append(chosenNumbers, number)
				}
			}
			if len(chosenNumbers) == 2 {
				gearRatio := 1
				for _, number := range chosenNumbers {
					gearRatio *= number.Value
				}
				sum += gearRatio
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
