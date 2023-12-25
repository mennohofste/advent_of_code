package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Steps     int
	Direction rune
}

const (
	SIMPLE  = 0
	COMPLEX = 1
)

func getDigPlan(complexity int) []Instruction {
	content, err := os.ReadFile("day_18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var instructions []Instruction
	for _, line := range strings.Split(strings.Trim(string(content), "\n"), "\n") {
		fields := strings.Fields(line)

		var instruction Instruction
		if complexity == SIMPLE {
			instruction.Direction = rune(fields[0][0])
			instruction.Steps, err = strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// (#4f4602) -> 4f460, 2
			color, direction := fields[2][2:len(fields[2])-2], fields[2][len(fields[2])-2]

			switch direction {
			case '0':
				instruction.Direction = 'R'
			case '1':
				instruction.Direction = 'D'
			case '2':
				instruction.Direction = 'L'
			case '3':
				instruction.Direction = 'U'
			}
			steps, err := strconv.ParseInt(color, 16, 0)
			if err != nil {
				log.Fatal(err)
			}
			instruction.Steps = int(steps)
		}

		instructions = append(instructions, instruction)
	}
	return instructions
}

type Position struct {
	Row int
	Col int
}

func getCoordinates(instructions []Instruction) ([]Position, int) {
	var positions []Position
	curPos := Position{0, 0}

	length := 0
	for _, instruction := range instructions {
		length += instruction.Steps

		switch instruction.Direction {
		case 'U':
			curPos.Row -= instruction.Steps
		case 'D':
			curPos.Row += instruction.Steps
		case 'L':
			curPos.Col -= instruction.Steps
		case 'R':
			curPos.Col += instruction.Steps
		}

		positions = append(positions, curPos)
	}

	return positions, length
}

func determinant(a, b Position) int {
	return a.Col*b.Row - a.Row*b.Col
}

func getTotalArea(coordinates []Position) int {
	// Shoelace formula
	total := 0
	for i := 0; i < len(coordinates); i++ {
		total += determinant(coordinates[i], coordinates[(i+1)%len(coordinates)])
	}
	return total / 2
}

func getTotalPoints(area, boundary int) int {
	// interior points = i
	// total area = A
	// boundary points = b
	//
	// Pick's theorem
	// A = i + b/2 - 1
	// -> i = A - b/2 + 1
	//
	// total_points = i + b
	//              = A - b/2 + 1 + b
	//              = A + b/2 + 1
	return area + boundary/2 + 1
}

func part1() int {
	instructions := getDigPlan(SIMPLE)
	coordinates, length := getCoordinates(instructions)
	area := getTotalArea(coordinates)
	return getTotalPoints(area, length)
}

func part2() int {
	instructions := getDigPlan(COMPLEX)
	coordinates, length := getCoordinates(instructions)
	area := getTotalArea(coordinates)
	return getTotalPoints(area, length)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
