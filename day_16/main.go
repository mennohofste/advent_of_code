package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Tile struct {
	Beams Direction
	Item  rune
}

func (t Tile) Energised() bool {
	return t.Beams > 0
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
			row = append(row, Tile{Item: item, Beams: Direction(0)})
		}
		contraption = append(contraption, row)
	}

	return contraption
}

type Direction byte

const (
	RIGHT Direction = 1 << iota
	DOWN
	LEFT
	UP
)

func (contraption Contraption) TraceBeam(row, col int, direction Direction) int {
	if row < 0 || row >= len(contraption) || col < 0 || col >= len(contraption[0]) || (contraption[row][col].Beams&direction) > 0 {
		return 0
	}

	energy := 0
	if !contraption[row][col].Energised() {
		energy++
	}
	contraption[row][col].Beams |= direction

	switch contraption[row][col].Item {
	case '|':
		if direction == RIGHT || direction == LEFT {
			energy += contraption.TraceBeam(row-1, col, UP)
			energy += contraption.TraceBeam(row+1, col, DOWN)
			return energy
		}
	case '-':
		if direction == DOWN || direction == UP {
			energy += contraption.TraceBeam(row, col-1, LEFT)
			energy += contraption.TraceBeam(row, col+1, RIGHT)
			return energy
		}
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

	switch direction {
	case RIGHT:
		col++
	case DOWN:
		row++
	case LEFT:
		col--
	case UP:
		row--
	}
	return energy + contraption.TraceBeam(row, col, direction)
}

func (contraption Contraption) Reset() {
	for i, line := range contraption {
		for j := range line {
			contraption[i][j].Beams = Direction(0)
		}
	}
}

func part1(contraption Contraption) int {
	defer contraption.Reset()
	return contraption.TraceBeam(0, 0, RIGHT)
}

func part2(contraption Contraption) int {
	var maxEnergised int
	for i := 0; i < len(contraption); i++ {
		energy := contraption.TraceBeam(i, 0, RIGHT)
		maxEnergised = max(maxEnergised, energy)
		contraption.Reset()
	}

	for i := 0; i < len(contraption); i++ {
		energy := contraption.TraceBeam(i, len(contraption)-1, LEFT)
		maxEnergised = max(maxEnergised, energy)
		contraption.Reset()
	}

	for i := 0; i < len(contraption[0]); i++ {
		energy := contraption.TraceBeam(0, i, DOWN)
		maxEnergised = max(maxEnergised, energy)
		contraption.Reset()
	}

	for i := 0; i < len(contraption[0]); i++ {
		energy := contraption.TraceBeam(len(contraption[0])-1, i, UP)
		maxEnergised = max(maxEnergised, energy)
		contraption.Reset()
	}

	return maxEnergised
}

func main() {
	startTime := time.Now()
	contraption := getContraption()
	fmt.Println("Part 1:", part1(contraption), time.Since(startTime))
	fmt.Println("Part 2:", part2(contraption), time.Since(startTime))
}
