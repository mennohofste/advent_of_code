package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Universe []string

func getLines() Universe {
	bytes, err := os.ReadFile("day_11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.Trim(string(bytes), "\n"), "\n")
}

func (u Universe) CheckOccupation() ([]bool, []bool) {
	occupiedRows := make([]bool, len(u))
	occupiedCols := make([]bool, len(u[0]))
	for i, line := range u {
		for j, space := range line {
			if space == '#' {
				occupiedRows[i] = true
				occupiedCols[j] = true
			}
		}
	}

	return occupiedRows, occupiedCols
}

type Galaxy struct {
	Row int
	Col int
}

func (u Universe) Galaxies(occupiedRows []bool, occupiedCols []bool, expansionFactor int) []Galaxy {
	var galaxies []Galaxy
	iOffset := 0
	for i, row := range u {
		jOffset := 0
		for j, space := range row {
			if space == '#' {
				galaxies = append(galaxies, Galaxy{Row: i + iOffset, Col: j + jOffset})
			}
			if !occupiedCols[j] {
				jOffset += expansionFactor - 1
			}
		}
		if !occupiedRows[i] {
			iOffset += expansionFactor - 1
		}
	}
	return galaxies
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1() int {
	universe := getLines()
	occupiedRows, occupiedCols := universe.CheckOccupation()
	galaxies := universe.Galaxies(occupiedRows, occupiedCols, 2)

	totalDistance := 0
	for i := range galaxies {
		for j := i; j < len(galaxies); j++ {
			totalDistance += abs(galaxies[i].Row-galaxies[j].Row) + abs(galaxies[i].Col-galaxies[j].Col)
		}
	}
	return totalDistance
}

func part2() int {
	universe := getLines()
	occupiedRows, occupiedCols := universe.CheckOccupation()
	galaxies := universe.Galaxies(occupiedRows, occupiedCols, 1000000)

	totalDistance := 0
	for i := range galaxies {
		for j := i; j < len(galaxies); j++ {
			totalDistance += abs(galaxies[i].Row-galaxies[j].Row) + abs(galaxies[i].Col-galaxies[j].Col)
		}
	}
	return totalDistance
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
