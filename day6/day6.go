package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Day6(filename string) {
	fmt.Println("Day 6")
	input6 := Input(filename)
	fmt.Println("Part 1:", part1(input6))
	fmt.Println("Part 2:", part2(input6))
}

func Widget(filename string) *fyne.Container {
	resultLabel := widget.NewLabel("")
	input := Input(filename)
	button1 := widget.NewButton("Part 1", func() {
		resultLabel.SetText("Result: " + strconv.Itoa(part1(input)))
	})

	button2 := widget.NewButton("Part 2", func() {
		resultLabel.SetText("Result: " + strconv.Itoa(part2(input)))
	})

	buttonRow := container.NewHBox(button1, button2)
	return container.NewVBox(buttonRow, resultLabel)
}

func part1(in []string) int {
	input := copyInput(in)
	dir := 0
	y, x := findStart(input)

	for inBound(input, x, y) {
		newline := []rune(input[y])
		newline[x] = 'X'
		input[y] = string(newline)
		x, y, dir = doLoop(x, y, dir, input)
	}

	count := 0
	for _, l := range input {
		for _, r := range l {
			if r == 'X' {
				count++
			}
		}
	}
	return count
}

func part2(input []string) int {
	count := 0
	for i, l := range input {
		for j := range l {
			if checkObstruction(input, j, i) {
				count++
			}
		}
	}
	return count
}

func checkObstruction(in []string, ox int, oy int) bool {
	input := copyInput(in)
	newline := []rune(input[oy])
	newline[ox] = '#'
	input[oy] = string(newline)
	dir := 0
	y, x := findStart(input)
	loop := 0

	for inBound(input, x, y) && loop < 99999 {
		x, y, dir = doLoop(x, y, dir, input)
		loop++
	}
	return loop >= 9999
}

func doLoop(x int, y int, dir int, input []string) (int, int, int) {
	yn := y + dirs[dir][0]
	xn := x + dirs[dir][1]

	h := len(input)
	w := len(input[0])

	if yn >= h || yn < 0 || xn >= w || xn < 0 {
		return xn, yn, dir
	} else if input[yn][xn] == '#' {
		dir++
		if dir == len(dirs) {
			dir = 0
		}
		return x, y, dir
	}
	return xn, yn, dir
}

func Input(filename string) []string {
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

func copyInput(input []string) []string {
	out := []string{}

	for _, l := range input {
		out = append(out, strings.Clone(l))
	}
	return out
}

func findStart(input []string) (int, int) {
	for i, v := range input {
		for j, w := range v {
			if string(w) == "^" {
				return i, j
			}
		}
	}
	return 0, 0
}

func inBound(input []string, x int, y int) bool {
	h := len(input)
	w := len(input[0])
	return x < w && x >= 0 && y < h && y >= 0
}
