package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func getLines() []string {
	bytes, err := os.ReadFile("day_05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.Trim(string(bytes), "\n"), "\n")
}

func getCatgeoryMaps(lines []string) [][]Mapping {
	var category_maps [][]Mapping
	var maps []Mapping
	for _, line := range lines[1:] {
		if line != "" && !strings.Contains(line, ":") {
			maps = append(maps, getMapping(line))
		}
		if line == "" && len(maps) != 0 {
			category_maps = append(category_maps, maps)
			maps = []Mapping{}
		}
	}
	return append(category_maps, maps)
}

func getLocation(seed uint32, category_maps [][]Mapping) uint32 {
	var done bool
	for _, maps := range category_maps {
		for _, m := range maps {
			seed, done = m.Map(seed)
			if done {
				break
			}
		}
	}
	return seed
}

func getSeeds(line string) []uint32 {
	var seeds []uint32
	for _, n := range strings.Fields(strings.Split(line, ":")[1]) {
		number, err := strconv.ParseUint(n, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, uint32(number))
	}
	return seeds
}

type Mapping struct {
	SourceStart uint32
	DestStart   uint32
	RangeLength uint32
}

func (m Mapping) Map(x uint32) (uint32, bool) {
	if x >= m.SourceStart && x < m.SourceStart+m.RangeLength {
		return x + m.DestStart - m.SourceStart, true
	}
	return x, false
}

func getMapping(line string) Mapping {
	var map_numbers [3]uint32
	for i, n := range strings.Fields(line) {
		number, err := strconv.ParseUint(n, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		map_numbers[i] = uint32(number)
	}

	return Mapping{
		DestStart:   map_numbers[0],
		SourceStart: map_numbers[1],
		RangeLength: map_numbers[2],
	}
}

func part1() uint32 {
	lines := getLines()
	seeds := getSeeds(lines[0])
	category_maps := getCatgeoryMaps(lines[1:])

	minLoc := getLocation(seeds[0], category_maps)
	var loc uint32
	for _, seed := range seeds {
		loc = getLocation(seed, category_maps)
		if loc < minLoc {
			minLoc = loc
		}
	}

	return minLoc
}

func worker(category_maps [][]Mapping, seed_ch <-chan uint32, result_ch chan<- uint32, wg *sync.WaitGroup) {
	defer wg.Done()
	for seed := range seed_ch {
		result_ch <- getLocation(seed, category_maps)
	}
}

func part2() uint32 {
	lines := getLines()
	seeds := getSeeds(lines[0])
	category_maps := getCatgeoryMaps(lines[1:])

	minLoc := getLocation(seeds[0], category_maps)
	var loc uint32
	for i := 0; i < len(seeds); i += 2 {
		for j := uint32(0); j < seeds[i+1]; j++ {
			loc = getLocation(seeds[i]+j, category_maps)
			if loc < minLoc {
				minLoc = loc
			}
		}
	}

	return minLoc
}

// func part2_multi() uint32 {
// 	lines := getLines()
// 	seeds := getSeeds(lines[0])
// 	category_maps := getCatgeoryMaps(lines[1:])
//
// 	var wg sync.WaitGroup
// 	seed_ch := make(chan uint32)
// 	result_ch := make(chan uint32)
// 	for w := 0; w < 1; w++ {
// 		wg.Add(1)
// 		go worker(category_maps, seed_ch, result_ch, &wg)
// 	}
//
// 	go func() {
// 		for i := 0; i < len(seeds); i += 2 {
// 			for j := uint32(0); j < seeds[i+1]; j++ {
// 				seed_ch <- seeds[i] + j
// 			}
// 		}
// 		close(seed_ch)
// 	}()
//
// 	go func() {
// 		wg.Wait()
// 		close(result_ch)
// 	}()
//
// 	minLoc := <-result_ch
// 	bar := progressbar.Default(int64(1800000000))
// 	count := 0
// 	for result := range result_ch {
// 		bar.Add(1)
// 		if result < minLoc {
// 			minLoc = result
// 		}
// 		if count == 100000000 {
// 			return minLoc
// 		}
// 		count++
// 	}
//
// 	return minLoc
// }

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
