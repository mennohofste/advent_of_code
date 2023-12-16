package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLine() []string {
	content, err := os.ReadFile("day_15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.ReplaceAll(string(content), "\n", ""), ",")
}

func hash(step string) int {
	currentValue := 0
	for _, c := range step {
		currentValue += int(c)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func part1() int {
	initSequence := getLine()

	sum := 0
	for _, step := range initSequence {
		sum += hash(step)
	}
	return sum
}

type Lens struct {
	Label       string
	FocalLength int
}

func removeLens(hashMap map[int][]Lens, label string) {
	lenses := hashMap[hash(label)]
	for i, lens := range lenses {
		if lens.Label == label {
			hashMap[hash(label)] = append(lenses[:i], lenses[i+1:]...)
			return
		}
	}
}

func setLens(hashMap map[int][]Lens, step string) {
	label, f, found := strings.Cut(step, "=")
	if !found {
		log.Fatal("Invalid input")
	}

	focalLength, err := strconv.Atoi(f)
	if err != nil {
		log.Fatal("Invalid input")
	}

	newLens := Lens{Label: label, FocalLength: focalLength}
	lenses := hashMap[hash(label)]
	for i, lens := range lenses {
		if lens.Label == label {
			lenses[i] = newLens
			hashMap[hash(label)] = lenses
			return
		}
	}
	hashMap[hash(label)] = append(hashMap[hash(label)], newLens)
}

func focussingPower(hashMap map[int][]Lens) int {
	sum := 0
	for box, lenses := range hashMap {
		for slot, lens := range lenses {
			sum += (1 + box) * (1 + slot) * lens.FocalLength
		}
	}
	return sum
}

func part2() int {
	hashMap := make(map[int][]Lens)
	initSequence := getLine()
	for _, step := range initSequence {
		if step[len(step)-1] == '-' {
			removeLens(hashMap, step[:len(step)-1])
		} else {
			setLens(hashMap, step)
		}
	}
	return focussingPower(hashMap)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
