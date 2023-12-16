package main

import (
	"log"
	"os"
	"strings"
)

func readLines() []string {
	content, err := os.ReadFile("day_08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content = content[:len(content)-1] // remove trailing newline
	return strings.Split(string(content), "\n")
}

type Node struct {
	Left  *Node
	Right *Node
	Name  string
}

func (n Node) String() string {
	return n.Name
}

func getNodes(lines []string) map[string]*Node {
	nodes := map[string]*Node{}
	for _, line := range lines {
		fields := strings.Fields(line)
		name := fields[0]
		nodes[name] = &Node{Name: name}
	}

	for _, line := range lines {
		fields := strings.Fields(line)
		name := fields[0]
		left := fields[2][1 : len(fields[2])-1]
		right := fields[3][:len(fields[3])-1]

		nodes[name].Left = nodes[left]
		nodes[name].Right = nodes[right]
	}

	return nodes
}

func part1() int {
	lines := readLines()
	instructions := lines[0]
	nodeMap := getNodes(lines[2:])

	node := *nodeMap["AAA"]
	NSteps := 0
	for node.Name != "ZZZ" {
		if instructions[NSteps%len(instructions)] == 'L' {
			node = *node.Left
		} else {
			node = *node.Right
		}
		NSteps++
	}
	return NSteps
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2() int {
	lines := readLines()
	instructions := lines[0]
	nodeMap := getNodes(lines[2:])

	var nodes []Node
	for _, node := range nodeMap {
		if node.Name[len(node.Name)-1] == 'A' {
			nodes = append(nodes, *node)
		}
	}

	var finishingIndices []int
	for _, node := range nodes {
		NSteps := 0
		for node.Name[len(node.Name)-1] != 'Z' {
			if instructions[NSteps%len(instructions)] == 'L' {
				node = *node.Left
			} else {
				node = *node.Right
			}
			NSteps++
		}

		finishingIndices = append(finishingIndices, NSteps)
	}

	return LCM(finishingIndices...)
}

func main() {
	println("Part 1:", part1())
	println("Part 2:", part2())
}
