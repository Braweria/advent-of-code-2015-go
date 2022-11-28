package days

import (
	"bufio"
	"fmt"
)

type Coord struct {
	x, y int
}

type Visited map[Coord]int

type Santa struct {
	position     Coord
	visited      Visited
	uniqueHouses int
}

func (santa *Santa) moveTo(value string) {
	switch value {
	case "^":
		santa.position.y++
	case "v":
		santa.position.y--
	case "<":
		santa.position.x--
	case ">":
		santa.position.x++
	}
}

func (santa *Santa) deliver() {
	visited, exists := santa.visited[santa.position]
	if exists {
		santa.visited[santa.position] = visited + 1
	} else {
		santa.visited[santa.position] = 1
		santa.uniqueHouses++
	}
}

func D03() {
	file := GetFile("inputs/03.txt")
	defer file.Close()

	start := Coord{0, 0}
	startVisited := Visited{start: 1}
	santa := Santa{start, startVisited, 1}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := scanner.Text()
		santa.moveTo(char)
		santa.deliver()
	}

	fmt.Println(santa.uniqueHouses)
}
