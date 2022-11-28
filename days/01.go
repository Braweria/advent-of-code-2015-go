package days

import (
	"bufio"
	"log"
	"os"
)

func D01() {
	solution("inputs/01.txt")
}

func solution(filename string) {
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
}
