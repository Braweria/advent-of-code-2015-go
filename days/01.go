package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func D01() {
	fmt.Println("Day 01")
	part1("inputs/01.txt")
}

func part1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	floor := 0
	basement := -1

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for i := 1; scanner.Scan(); i++ {
		char := scanner.Text()
		if char == "(" {
			floor++
		}
		if char == ")" {
			floor--
			if basement == -1 && floor == -1 {
				basement = i
			}
		}
	}
	fmt.Printf("Part I:  %v\nPart II: %v", floor, basement)
}
