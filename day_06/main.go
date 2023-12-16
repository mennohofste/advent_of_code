package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLines() []string {
	content, err := os.ReadFile("day_06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.Trim(string(content), "\n"), "\n")
}

type Race struct {
	Time     int
	Distance int
}

func (r Race) NumberOfWaysToWin() int {
	numberOfWins := 0
	for time := 1; time < r.Time; time++ {
		if (r.Time-time)*time > r.Distance {
			numberOfWins++
		}
	}
	return numberOfWins
}

func part1() int {
	lines := getLines()
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]

	totalProd := 1
	for i := range times {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatal(err)
		}

		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatal(err)
		}

		race := Race{Time: time, Distance: distance}
		totalProd *= race.NumberOfWaysToWin()
	}

	return totalProd
}

func part2() int {
	lines := getLines()

	timeRaw := ""
	for _, t := range strings.Fields(lines[0])[1:] {
		timeRaw += t
	}
	time, err := strconv.Atoi(timeRaw)
	if err != nil {
		log.Fatal(err)
	}

	distanceRaw := ""
	for _, t := range strings.Fields(lines[1])[1:] {
		distanceRaw += t
	}
	distance, err := strconv.Atoi(distanceRaw)
	if err != nil {
		log.Fatal(err)
	}

	race := Race{Time: time, Distance: distance}
	return race.NumberOfWaysToWin()
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
