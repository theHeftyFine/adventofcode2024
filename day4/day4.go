package day4

import (
	"bufio"
	"log"
	"os"
	"slices"

	"fyne.io/fyne/v2"
	daydisplay "github.com/theheftyfine/adventofcode2024/display"
)

type day struct{}

var display = daydisplay.BasicDisplay[[]string]{
	DayRunner: day{},
}

func Display(filename string) *fyne.Container {
	return display.Widget(filename)
}

func (day) Part1(input []string, cont *fyne.Container) int {
	horizontal := horizontalCount(input)

	vertical := verticalCount(input)

	diagonalBottomTop := diagonalCount(input)

	diagonalTopBottom := diagonalCount2(input)

	return horizontal + vertical + diagonalBottomTop + diagonalTopBottom
}

func (day) Part2(input []string, cont *fyne.Container) int {
	var count = 0
	for i := 0; i < len(input)-2; i++ {
		for j := 0; j < len(input[0])-2; j++ {
			mid := string(input[i+1][j+1])
			tl := string(input[i][j])
			bl := string(input[i+2][j])
			tr := string(input[i][j+2])
			br := string(input[i+2][j+2])

			if mid == "A" && ((tl == "M") && (br == "S") || (tl == "S" && br == "M")) && ((bl == "M") && (tr == "S") || (bl == "S" && tr == "M")) {
				count++
			}
		}
	}
	return count
}

func (day) Input(filename string) []string {
	var lines = []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func countChristmasFromLine(line string) int {
	var count = 0
	var i = 0
	for i < len(line)-3 {
		sub := line[i : i+4]
		if sub == "XMAS" || sub == "SAMX" {
			count++
		}
		i++
	}
	return count
}

func horizontalCount(input []string) int {
	var count = 0
	for _, line := range input {
		count += countChristmasFromLine(line)
	}
	return count
}

func verticalCount(input []string) int {
	var count = 0
	for i := 0; i < len(input)-3; i++ {
		cols := input[i : i+4]
		for j := 0; j < len(cols[0]); j++ {
			vertsub := string(cols[0][j]) + string(cols[1][j]) + string(cols[2][j]) + string(cols[3][j])
			if vertsub == "XMAS" || vertsub == "SAMX" {
				count++
			}
		}
	}
	return count
}

func diagonalCount(input []string) int {
	var count = 0
	var left = []string{}
	var right = []string{}
	// diagonal
	for i := 0; i < len(input); i++ {
		// left side diagonal
		// x x x x o
		// x x x o o
		// x x o o o
		// x o o o o
		// o o o o o
		var diagonalLeft = ""
		// right side diagonal
		// o o o o o
		// o o o o x
		// o o o x x
		// o o x x x
		// o x x x x
		var diagonalRight = ""

		for x := i; x > -1; x-- {
			xOpp := (len(input) - 1) - x
			yOpp := (len(input) - 1) - (i - x)
			diagonalLeft += string(input[x][i-x])
			diagonalRight += string(input[xOpp][yOpp])
		}
		left = append(left, diagonalLeft)
		right = append(right, diagonalRight)
	}

	slices.Reverse(right)
	con := slices.Concat(left, right[1:])
	for _, v := range con {
		count += countChristmasFromLine(v)
	}
	return count
}

func diagonalCount2(input []string) int {
	var count = 0
	var left = []string{}
	var right = []string{}
	// diagonal
	for i := 0; i < len(input); i++ {
		// left side diagonal
		// o x x x x
		// o o x x x
		// o o o x x
		// o o o o x
		// o o o o o
		var diagonalLeft = ""
		//  side diagonal
		// o o o o o
		// x o o o o
		// x x o o o
		// x x x o o
		// x x x x o
		var diagonalRight = ""
		for x := i; x > -1; x-- {
			xOpp := (len(input) - 1) - x
			yOpp := (i - x)
			diagonalLeft += string(input[x][len(input)-1-(i-x)])
			diagonalRight += string(input[xOpp][yOpp])
		}
		left = append(left, diagonalLeft)
		right = append(right, diagonalRight)
	}

	// middle diagonal
	// x o o o o
	// o x o o o
	// o o x o o
	// o o o x o
	// o o o o x
	slices.Reverse(right)
	con := slices.Concat(left, right[1:])
	for _, v := range con {
		c := countChristmasFromLine(v)
		count += c
	}
	return count
}
