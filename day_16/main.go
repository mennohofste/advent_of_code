package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Tile struct {
	Beams map[Direction]bool
	Item  rune
}

func (t Tile) Energised() bool {
	for _, energised := range t.Beams {
		if energised {
			return true
		}
	}
	return false
}

type Contraption [][]Tile

func getContraption() Contraption {
	content, err := os.ReadFile("day_16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var contraption [][]Tile
	for _, line := range strings.Split(strings.Trim(string(content), "\n"), "\n") {
		var row []Tile
		for _, item := range line {
			row = append(row, Tile{Item: item, Beams: make(map[Direction]bool)})
		}
		contraption = append(contraption, row)
	}

	return contraption
}

type Direction int

const (
	RIGHT Direction = iota
	DOWN
	LEFT
	UP
)

func newCoordinates(row, col int, direction Direction) (int, int) {
	switch direction {
	case RIGHT:
		return row, col + 1
	case DOWN:
		return row + 1, col
	case LEFT:
		return row, col - 1
	case UP:
		return row - 1, col
	}
	return 0, 0
}

func (contraption Contraption) TraceBeam(row, col int, direction Direction) {
	if row < 0 || row >= len(contraption) || col < 0 || col >= len(contraption[0]) || contraption[row][col].Beams[direction] {
		return
	}
	contraption[row][col].Beams[direction] = true

	if contraption[row][col].Item == '|' && (direction == RIGHT || direction == LEFT) {
		contraption.TraceBeam(row-1, col, UP)
		contraption.TraceBeam(row+1, col, DOWN)
		return
	}
	if contraption[row][col].Item == '-' && (direction == DOWN || direction == UP) {
		contraption.TraceBeam(row, col-1, LEFT)
		contraption.TraceBeam(row, col+1, RIGHT)
		return
	}

	switch contraption[row][col].Item {
	case '/':
		switch direction {
		case RIGHT:
			direction = UP
		case DOWN:
			direction = LEFT
		case LEFT:
			direction = DOWN
		case UP:
			direction = RIGHT
		}
	case '\\':
		switch direction {
		case RIGHT:
			direction = DOWN
		case DOWN:
			direction = RIGHT
		case LEFT:
			direction = UP
		case UP:
			direction = LEFT
		}
	}

	newRow, newCol := newCoordinates(row, col, direction)
	contraption.TraceBeam(newRow, newCol, direction)
}

func (contraption Contraption) Energised() int {
	sum := 0
	for _, line := range contraption {
		for _, tile := range line {
			if tile.Energised() {
				sum++
			}
		}
	}
	return sum
}

func (contraption Contraption) Reset() {
	for _, line := range contraption {
		for _, tile := range line {
			for direction := range tile.Beams {
				tile.Beams[direction] = false
			}
		}
	}
}

func part1() int {
	contraption := getContraption()
	contraption.TraceBeam(0, 0, RIGHT)
	return contraption.Energised()
}

func part2() int {
	contraption := getContraption()

	var maxEnergised int
	for i := 0; i < len(contraption); i++ {
		contraption.TraceBeam(i, 0, RIGHT)
		maxEnergised = max(maxEnergised, contraption.Energised())
		contraption.Reset()
	}

	for i := 0; i < len(contraption); i++ {
		contraption.TraceBeam(i, len(contraption)-1, LEFT)
		maxEnergised = max(maxEnergised, contraption.Energised())
		contraption.Reset()
	}

	for i := 0; i < len(contraption[0]); i++ {
		contraption.TraceBeam(0, i, DOWN)
		maxEnergised = max(maxEnergised, contraption.Energised())
		contraption.Reset()
	}

	for i := 0; i < len(contraption[0]); i++ {
		contraption.TraceBeam(len(contraption[0])-1, i, UP)
		maxEnergised = max(maxEnergised, contraption.Energised())
		contraption.Reset()
	}

	return maxEnergised
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
