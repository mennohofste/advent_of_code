package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getLines() []string {
	content, err := os.ReadFile("day_10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.Trim(string(content), "\n"), "\n")
}

func getToken(char rune, row, col int) Token {
	token := Token{Name: char, Distance: -1, Row: row, Col: col}

	switch char {
	case 'S':
		token.ConnectsNorth = true
		token.ConnectsEast = true
		token.ConnectsSouth = true
		token.ConnectsWest = true
		token.Distance = 0
	case '|':
		token.ConnectsNorth = true
		token.ConnectsSouth = true
	case '-':
		token.ConnectsEast = true
		token.ConnectsWest = true
	case 'L':
		token.ConnectsNorth = true
		token.ConnectsEast = true
	case 'J':
		token.ConnectsNorth = true
		token.ConnectsWest = true
	case '7':
		token.ConnectsSouth = true
		token.ConnectsWest = true
	case 'F':
		token.ConnectsEast = true
		token.ConnectsSouth = true
	}
	return token
}

type Token struct {
	Distance      int
	Row           int
	Col           int
	Name          rune
	ConnectsNorth bool
	ConnectsEast  bool
	ConnectsSouth bool
	ConnectsWest  bool
}

func (t Token) ConnectsTo(o *Token) bool {
	switch {
	case t.Row == o.Row && t.Col == o.Col-1:
		return t.ConnectsEast && o.ConnectsWest
	case t.Row == o.Row && t.Col == o.Col+1:
		return t.ConnectsWest && o.ConnectsEast
	case t.Row == o.Row-1 && t.Col == o.Col:
		return t.ConnectsSouth && o.ConnectsNorth
	case t.Row == o.Row+1 && t.Col == o.Col:
		return t.ConnectsNorth && o.ConnectsSouth
	default:
		return false
	}
}

func (t Token) IsChanged() bool {
	return t.Distance >= 0
}

func (t *Token) InferToken(field Field) {
	// above
	above := field[t.Row-1][t.Col]
	below := field[t.Row+1][t.Col]
	left := field[t.Row][t.Col-1]
	right := field[t.Row][t.Col+1]
	switch {
	case above.ConnectsSouth && left.ConnectsEast:
		t.Name = 'J'
	case above.ConnectsSouth && right.ConnectsWest:
		t.Name = 'L'
	case above.ConnectsSouth && below.ConnectsNorth:
		t.Name = '|'
	case below.ConnectsNorth && left.ConnectsEast:
		t.Name = '7'
	case below.ConnectsNorth && right.ConnectsWest:
		t.Name = 'F'
	case left.ConnectsEast && right.ConnectsWest:
		t.Name = '-'
	default:
		log.Fatal(above.ConnectsSouth, right.ConnectsWest, below.ConnectsNorth, left.ConnectsEast)
	}
}

func (t Token) String() string {
	return string(t.Name)
}

type Field [][]*Token

func tokenize(lines []string) (Field, *Token) {
	var startToken *Token
	var field Field
	for i, line := range lines {
		var tokenRow []*Token
		for j, char := range line {
			token := getToken(char, i, j)
			tokenRow = append(tokenRow, &token)
			if char == 'S' {
				startToken = &token
			}
		}
		field = append(field, tokenRow)
	}
	return field, startToken
}

func (f Field) String() string {
	field := ""
	for _, line := range field {
		field += fmt.Sprintln(line)
	}
	return field
}

func (field Field) TraceRoute(startToken *Token) int {
	changedTokens := []*Token{startToken}
	var markedForChange []*Token
	for {
		for _, token := range changedTokens {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if token.Row+i >= 0 && token.Row+i < len(field) && token.Col+j >= 0 && token.Col+j < len(field[0]) {
						newToken := field[token.Row+i][token.Col+j]
						if !newToken.IsChanged() && newToken.ConnectsTo(token) {
							markedForChange = append(markedForChange, newToken)
						}
					}
				}
			}
		}

		if len(markedForChange) == 0 {
			break
		}

		for _, token := range markedForChange {
			token.Distance = changedTokens[0].Distance + 1
		}
		changedTokens = markedForChange
		markedForChange = []*Token{}
	}
	return changedTokens[0].Distance
}

func (field Field) MarkInner() int {
	innerCount := 0
	for _, row := range field {
		inner := false
		var startBoundaryToken *Token
		for _, token := range row {
			if inner && token.Name == '.' {
				innerCount++
				token.Name = 'I'
			}
			if token.Name == '|' {
				inner = !inner
			}
			if token.Name == 'L' || token.Name == 'F' {
				startBoundaryToken = token
			}
			if startBoundaryToken != nil && startBoundaryToken.Name == 'L' && token.Name == '7' {
				inner = !inner
			}
			if startBoundaryToken != nil && startBoundaryToken.Name == 'F' && token.Name == 'J' {
				inner = !inner
			}
		}
	}
	return innerCount
}

func (field Field) DowngradeUnusedPipes() {
	for _, row := range field {
		for _, token := range row {
			if token.Distance < 0 {
				token.Name = '.'
			}
		}
	}
}

func part1() int {
	lines := getLines()
	field, startToken := tokenize(lines)
	maxDistance := field.TraceRoute(startToken)
	return maxDistance
}

func part2() int {
	lines := getLines()
	field, startToken := tokenize(lines)
	field.TraceRoute(startToken)
	field.DowngradeUnusedPipes()
	startToken.InferToken(field)
	return field.MarkInner()
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
