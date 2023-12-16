package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readLines() []string {
	content, err := os.ReadFile("day_09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content = content[:len(content)-1]
	return strings.Split(string(content), "\n")
}

func getSensorData(lines []string) [][]int {
	var sensorData [][]int
	for _, line := range lines {
		var history []int
		for _, value := range strings.Fields(line) {
			if v, err := strconv.Atoi(value); err == nil {
				history = append(history, v)
			}
		}
		sensorData = append(sensorData, history)
	}
	return sensorData
}

func getExtrapolatedSensorData(sensorData [][]int) [][][]int {
	extrapolatedSensorData := make([][][]int, len(sensorData))
	for i, dataRow := range sensorData {
		extrapolatedSensorData[i] = append(extrapolatedSensorData[i], dataRow)

		done := false
		for !done {
			done = true
			oldN := dataRow[0]
			currentStatus := make([]int, len(dataRow)-1)
			for i, v := range dataRow[1:] {
				currentStatus[i] = v - oldN
				oldN = v
				if currentStatus[i] != 0 {
					done = false
				}
			}
			extrapolatedSensorData[i] = append(extrapolatedSensorData[i], currentStatus)
			dataRow = currentStatus
		}
	}
	return extrapolatedSensorData
}

func part1() int {
	lines := readLines()
	sensorData := getSensorData(lines)
	extrapolatedSensorData := getExtrapolatedSensorData(sensorData)

	sum := 0
	for _, history := range extrapolatedSensorData {
		for _, status := range history {
			sum += status[len(status)-1]
		}
	}

	return sum
}

func part2() int {
	lines := readLines()
	sensorData := getSensorData(lines)
	extraPolatedHistory := getExtrapolatedSensorData(sensorData)

	sum := 0
	for _, history := range extraPolatedHistory {
		lastVal := 0
		slices.Reverse(history)
		for _, status := range history {
			lastVal = status[0] - lastVal
		}
		sum += lastVal
	}

	return sum
}

func main() {
	println("Part 1:", part1())
	println("Part 2:", part2())
}
