package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strings"
)

func getMap() (map[Position]int, Position) {
	content, err := os.ReadFile("day_17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	islandMap := make(map[Position]int)
	var end Position
	for r, line := range strings.Fields(string(content)) {
		for c, char := range line {
			end = Position{Row: r, Col: c}
			islandMap[end] = int(char - '0')
		}
	}

	return islandMap, end
}

type Position struct {
	Row int
	Col int
}

func (p Position) Add(d Direction) Position {
	p.Row += d.Row
	p.Col += d.Col
	return p
}

type Direction Position

func (d Direction) Add(other Direction) Direction {
	d.Row += other.Row
	d.Col += other.Col
	return d
}

func (d Direction) Mul(n int) Direction {
	d.Row *= n
	d.Col *= n
	return d
}

var (
	NORTH = Direction{Row: -1, Col: 0}
	EAST  = Direction{Row: 0, Col: 1}
	SOUTH = Direction{Row: 1, Col: 0}
	WEST  = Direction{Row: 0, Col: -1}
)

type State struct {
	Position  Position
	Direction Direction
}

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PQ[T]) GPush(v T, p int)  { heap.Push(q, pqi[T]{v, p}) }
func (q *PQ[T]) GPop() (T, int)    { x := heap.Pop(q).(pqi[T]); return x.v, x.p }

func uniformCostSearch(islandMap map[Position]int, end Position, minMoves, maxMoves int) int {
	var frontier PQ[State]
	frontier.GPush(State{Direction: EAST}, 0)
	frontier.GPush(State{Direction: SOUTH}, 0)

	seen := make(map[State]interface{})

	for {
		state, cost := frontier.GPop()
		if state.Position == end {
			return cost
		}

		if _, ok := seen[state]; ok {
			continue
		}
		seen[state] = true

		for moves := minMoves; moves <= maxMoves; moves++ {
			newPosition := state.Position.Add(state.Direction.Mul(moves))

			// Do not go out of bounds
			if _, ok := islandMap[newPosition]; !ok {
				continue
			}

			newCost := cost
			for j := 1; j <= moves; j++ {
				newCost += islandMap[state.Position.Add(state.Direction.Mul(j))]
			}

			switch state.Direction {
			case NORTH, SOUTH:
				frontier.GPush(State{newPosition, EAST}, newCost)
				frontier.GPush(State{newPosition, WEST}, newCost)
			case EAST, WEST:
				frontier.GPush(State{newPosition, NORTH}, newCost)
				frontier.GPush(State{newPosition, SOUTH}, newCost)
			}
		}
	}
}

func part1() int {
	islandMap, end := getMap()
	return uniformCostSearch(islandMap, end, 1, 3)
}

func part2() int {
	islandMap, end := getMap()
	return uniformCostSearch(islandMap, end, 4, 10)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
