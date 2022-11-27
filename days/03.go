package days

import "fmt"

type Coord struct {
	x, y int
}

func (coord *Coord) update(value string) {
	switch value {
	case "^":
		coord.y++
	case "v":
		coord.y--
	case "<":
		coord.x--
	case ">":
		coord.x++
	}
}

func D03() {
	file := GetFile("inputs/03.txt")
	defer file.Close()

	coords := map[Coord]int{Coord{0, 0}: 0}

	coord := &Coord{0, 0}

	fmt.Println(x, y)
}