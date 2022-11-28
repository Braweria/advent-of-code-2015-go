package days

import (
	"bufio"
)

type Coord struct {
	x, y int
}

type Visited map[Coord]int

type Santa struct {
	position     Coord
	roboPosition Coord
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

func (santa *Santa) moveRoboTo(value string) {
	switch value {
	case "^":
		santa.roboPosition.y++
	case "v":
		santa.roboPosition.y--
	case "<":
		santa.roboPosition.x--
	case ">":
		santa.roboPosition.x++
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

func (santa *Santa) roboDeliver() {
	visited, exists := santa.visited[santa.roboPosition]
	if exists {
		santa.visited[santa.roboPosition] = visited + 1
	} else {
		santa.visited[santa.roboPosition] = 1
		santa.uniqueHouses++
	}
}

func D03() {
	file := GetFile("inputs/03.txt")
	defer file.Close()

	start := Coord{0, 0}
	startVisited := Visited{start: 2}
	santa := Santa{start, start, startVisited, 1}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for i := 1; scanner.Scan(); i++ {
		char := scanner.Text()
		if i%2 == 0 {
			santa.moveRoboTo(char)
			santa.roboDeliver()
		} else {
			santa.moveTo(char)
			santa.deliver()
		}
	}
}
