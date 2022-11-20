package days

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func D02() {
	fmt.Printf("\nDay 02\n")
	solution2("inputs/02.txt")
}

func solution2(filename string) {
	file := GetFile(filename)

	var wrappingPaper float64
	wrappingPaper = -1

	var prisms []PrismSides

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		stringSizes := strings.Split(scanner.Text(), "x")
		var dimensions []float64
		for _, arg := range stringSizes {
			if num, err := strconv.ParseFloat(arg, 64); err == nil {
				dimensions = append(dimensions, num)
			}
		}
		prism := PrismSides{dimensions[0], dimensions[1], dimensions[2]}
		prisms = append(prisms, prism)
	}

	fmt.Println(prisms)
	for _, prism := range prisms {
		paperArea := prism.paperArea()
		wrappingPaper += paperArea
	}

	fmt.Printf("Part I:  %v\nPart II: %v", wrappingPaper, 0)
}

type PrismSides struct {
	length, width, height float64
}

func (prism *PrismSides) paperArea() float64 {
	var lw, lh, wh float64

	lw = prism.length * prism.width * 2
	lh = prism.length * prism.height * 2
	wh = prism.width * prism.height * 2

	smallestArea := math.Min(lw, math.Min(lh, wh)) / 2

	return lw + lh + wh + smallestArea
}
