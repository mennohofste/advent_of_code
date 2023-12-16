package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Reflector struct {
	Field [][]rune
}

func (r Reflector) String() string {
	out := ""
	for _, line := range r.Field {
		out += fmt.Sprintf("%s\n", string(line))
	}
	return out
}

func (r *Reflector) Tilt(direction Direction) {
	var obstacles []int
	switch direction {
	case NORTH:
		obstacles = make([]int, len(r.Field[0]))
	case EAST:
		obstacles = make([]int, len(r.Field))
		for i := range obstacles {
			obstacles[i] = len(r.Field[i]) - 1
		}
	case SOUTH:
		obstacles = make([]int, len(r.Field[0]))
		for i := range obstacles {
			obstacles[i] = len(r.Field) - 1
		}
	case WEST:
		obstacles = make([]int, len(r.Field))
	}

	newField := make([][]rune, len(r.Field))
	for i := range newField {
		newField[i] = make([]rune, len(r.Field[i]))
		for j := range newField[i] {
			newField[i][j] = '.'
		}
	}

	for i, row := range r.Field {
		if direction == SOUTH {
			i = len(r.Field) - i - 1
			row = r.Field[i]
		}

		for j, spot := range row {
			if direction == EAST {
				j = len(row) - j - 1
				spot = row[j]
			}

			switch spot {
			case 'O':
				switch direction {
				case NORTH:
					newField[obstacles[j]][j] = spot
					obstacles[j]++
				case EAST:
					newField[i][obstacles[i]] = spot
					obstacles[i]--
				case SOUTH:
					newField[obstacles[j]][j] = spot
					obstacles[j]--
				case WEST:
					newField[i][obstacles[i]] = spot
					obstacles[i]++
				}
			case '#':
				newField[i][j] = spot
				switch direction {
				case NORTH:
					obstacles[j] = i + 1
				case EAST:
					obstacles[i] = j - 1
				case SOUTH:
					obstacles[j] = i - 1
				case WEST:
					obstacles[i] = j + 1
				}
			}
		}
	}

	r.Field = newField
}

func (r Reflector) Load() int {
	loadSum := 0
	maxLoad := len(r.Field)
	for i, row := range r.Field {
		for _, spot := range row {
			if spot == 'O' {
				loadSum += maxLoad - i
			}
		}
	}
	return loadSum
}

func (r *Reflector) Cycle() {
	r.Tilt(NORTH)
	r.Tilt(WEST)
	r.Tilt(SOUTH)
	r.Tilt(EAST)
}

func (r Reflector) IsEqual(other Reflector) bool {
	for i, row := range r.Field {
		for j, spot := range row {
			if spot != other.Field[i][j] {
				return false
			}
		}
	}
	return true
}

func getLines() [][]rune {
	content, err := os.ReadFile("day_14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var field [][]rune
	for _, line := range strings.Split(strings.Trim(string(content), "\n"), "\n") {
		field = append(field, []rune(line))
	}
	return field
}

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func part1() int {
	reflector := Reflector{Field: getLines()}
	reflector.Tilt(NORTH)
	return reflector.Load()
}

func (r *Reflector) SmartCycle(number int) {
	var cycleStart int
	var cycleLength int
	var previousReflectors []Reflector
	for {
		oldReflector := Reflector{Field: make([][]rune, len(r.Field))}
		copy(oldReflector.Field, r.Field)
		previousReflectors = append(previousReflectors, oldReflector)

		r.Cycle()

		for i, oldReflector := range previousReflectors {
			if r.IsEqual(oldReflector) {
				cycleStart = i
				cycleLength = len(previousReflectors) - i
				break
			}
		}
		if cycleLength > 0 {
			r.Field = previousReflectors[(number-cycleStart)%cycleLength+cycleStart].Field
			return
		}
	}
}

func part2() int {
	reflector := Reflector{Field: getLines()}
	reflector.SmartCycle(1000000000)
	return reflector.Load()
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
