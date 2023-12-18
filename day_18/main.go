package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Color     string
	Steps     int
	Direction rune
}

func getDigPlan() []Instruction {
	content, _ := os.ReadFile("day_18/input.txt")

	var instructions []Instruction
	for _, line := range strings.Split(strings.Trim(string(content), "\n"), "\n") {
		fields := strings.Fields(line)
		direction := rune(fields[0][0])
		steps, _ := strconv.Atoi(fields[1])
		color := fields[2]

		instructions = append(instructions, Instruction{Color: color, Steps: steps, Direction: direction})
	}
	return instructions
}

type Position struct {
	Row int
	Col int
}

func (p Position) Sub(other Position) Position {
	return Position{p.Row - other.Row, p.Col - other.Col}
}

func (p Position) Mul(factor int) Position {
	return Position{p.Row * factor, p.Col * factor}
}

func getMaxDimensions(instructions []Instruction) (Position, Position) {
	minPos := Position{0, 0}
	maxPos := Position{0, 0}
	curPos := Position{0, 0}

	for _, instruction := range instructions {
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

		if curPos.Row < minPos.Row {
			minPos.Row = curPos.Row
		}
		if curPos.Row > maxPos.Row {
			maxPos.Row = curPos.Row
		}
		if curPos.Col < minPos.Col {
			minPos.Col = curPos.Col
		}
		if curPos.Col > maxPos.Col {
			maxPos.Col = curPos.Col
		}
	}
	return maxPos.Sub(minPos), minPos.Mul(-1)
}

type CellState byte

func (c CellState) String() string {
	switch c {
	case UNKNOWN:
		return "?"
	case INSIDE:
		return "#"
	case OUTSIDE:
		return "."
	}
	return "E"
}

const (
	UNKNOWN CellState = 1<<iota - 1
	INSIDE
	OUTSIDE
)

func getGrid(dimensions Position, startPos Position, instructions []Instruction) [][]CellState {
	grid := make([][]CellState, dimensions.Row+1)
	for i := range grid {
		grid[i] = make([]CellState, dimensions.Col+1)
	}

	grid[startPos.Row][startPos.Col] = INSIDE
	for _, instruction := range instructions {
		for i := 0; i < instruction.Steps; i++ {
			switch instruction.Direction {
			case 'U':
				startPos.Row--
			case 'D':
				startPos.Row++
			case 'L':
				startPos.Col--
			case 'R':
				startPos.Col++
			}

			grid[startPos.Row][startPos.Col] = INSIDE
		}
	}

	return grid
}

func colorOutside(grid [][]CellState, pos Position) {
	if pos.Row < 0 || pos.Row >= len(grid) || pos.Col < 0 || pos.Col >= len(grid[pos.Row]) || grid[pos.Row][pos.Col] != UNKNOWN {
		return
	}

	grid[pos.Row][pos.Col] = OUTSIDE
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			colorOutside(grid, Position{pos.Row + i, pos.Col + j})
		}
	}
}

func colorGrid(grid [][]CellState) {
	for i := range grid {
		colorOutside(grid, Position{i, 0})
		colorOutside(grid, Position{i, len(grid[i]) - 1})
	}

	for i := range grid[0] {
		colorOutside(grid, Position{0, i})
		colorOutside(grid, Position{len(grid) - 1, i})
	}
}

func countInside(grid [][]CellState) int {
	total := 0
	for _, line := range grid {
		for _, cell := range line {
			if cell == OUTSIDE {
				total++
			}
		}
	}
	return len(grid)*len(grid[0]) - total
}

func part1() int {
	instructions := getDigPlan()
	dimensions, startPos := getMaxDimensions(instructions)
	grid := getGrid(dimensions, startPos, instructions)
	colorGrid(grid)
	return countInside(grid)
}

func part2() int {
	return 0
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
