package day8

import (
	"bufio"
	"github.com/theheftyfine/adventofcode2024/model"
	"log"
	"os"
	"slices"
	"strings"
)

type day struct {
	input []string
}

func (d day) part1() int {
	return part1(d.input)
}

func (d day) part2() int {
	return part2(d.input)
}

func (d day) Parts() []func() int {
	return []func() int{d.part1, d.part2}
}

func NewDay(filename string) model.DayRunner {
	return day{input: input(filename)}
}

func part1(input []string) int {
	return calcNodes(input, drawPart1)
}

func part2(input []string) int {
	return calcNodes(input, drawPart2)
}

func input(filename string) []string {
	out := []string{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out
}

func calcNodes(in []string, drawFunc func([]string, []int, []int)) int {
	input := copyInput(in)

	coordinates := getCoordinates(input)

	for _, v := range coordinates {
		coord := v[0]
		coords := v[1:]
		for len(coords) > 0 {
			for _, c := range coords {
				distA, distB := distance(coord, c)

				drawFunc(input, coord, distA)

				drawFunc(input, c, distB)
			}
			coord = coords[0]
			coords = coords[1:]
		}
	}

	sum := 0
	for _, l := range input {
		for _, r := range l {
			if r == '#' {
				sum++
			}
		}
	}
	return sum
}

func drawPart1(input []string, coord []int, dist []int) {
	node := slices.Clone(coord)
	node[0] += dist[0]
	node[1] += dist[1]

	if inBound(input, node[1], node[0]) {
		replace(input, node[0], node[1], '#')
	}
}

func drawPart2(input []string, coord []int, dist []int) {
	node := slices.Clone(coord)

	for inBound(input, node[1], node[0]) {
		replace(input, node[0], node[1], '#')
		node[0] += dist[0]
		node[1] += dist[1]
	}
}

func getCoordinates(input []string) map[rune][][]int {
	coordinates := make(map[rune][][]int)
	for i, line := range input {
		for j, r := range line {
			if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
				coordinates[r] = append(coordinates[r], []int{i, j})
			}
		}
	}

	return coordinates
}

func distance(coordA []int, coordB []int) ([]int, []int) {
	distV := coordA[0] - coordB[0]
	if distV < 0 {
		distV = distV * -1
	}

	distH := coordA[1] - coordB[1]
	if distH < 0 {
		distH = distH * -1
	}

	vecA := []int{0, 0}
	vecB := []int{0, 0}

	if coordA[0] < coordB[0] {
		vecA[0] = -distV
		vecB[0] = distV
	} else if coordA[0] > coordB[0] {
		vecA[0] = distV
		vecB[0] = -distV
	}

	if coordA[1] < coordB[1] {
		vecA[1] = -distH
		vecB[1] = distH
	} else if coordA[1] > coordB[1] {
		vecA[1] = distH
		vecB[1] = -distH
	}

	return vecA, vecB
}

func copyInput(input []string) []string {
	out := []string{}

	for _, l := range input {
		out = append(out, strings.Clone(l))
	}
	return out
}

func inBound(input []string, x int, y int) bool {
	h := len(input)
	w := len(input[0])
	return x < w && x >= 0 && y < h && y >= 0
}

func replace(input []string, y int, x int, r rune) {
	row := []rune(input[y])
	row[x] = r
	input[y] = string(row)
}
