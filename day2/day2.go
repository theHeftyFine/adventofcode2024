package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Day2(filename string) {
	fmt.Println("Day 2:")
	input2 := Input(filename)
	fmt.Println("Part 1:", part1(input2))
	fmt.Println("Part 2:", part2(input2))
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

func part1(input [][]int) int {
	var safe = 0

	for _, row := range input {
		if checkRow(row) {
			safe = safe + 1
		}
	}

	return safe
}

func part2(input [][]int) int {
	var safe = 0

	for _, row := range input {
		split := splitLevels(row)
		count := 1
		for _, variation := range split {
			if checkRow(variation) {
				safe++
				break
			}
			count++
		}
	}

	return safe
}

func checkRow(row []int) bool {
	var rising = 0
	var unsaferise = false
	var zeros = false

	for i, v := range row {

		if i < len(row)-1 {
			d := row[i+1] - v
			// no zero changes
			if d == 0 {
				zeros = true
				break
			}
			// check if the change is in acceptable levels
			if d > 3 || d < -3 {
				unsaferise = true
				break
			}
			// count rising changes only
			if d > 0 {
				rising++
			}
		}

	}
	return !zeros && !unsaferise && (rising == 0 || rising == len(row)-1)
}

func splitLevels(levels []int) [][]int {
	out := [][]int{levels}
	for i := 0; i < len(levels); i++ {
		var o = levels[:i]
		if i < len(levels)-1 {
			o = slices.Concat(o, levels[i+1:])
		}
		out = append(out, o)
	}
	return out
}

func Input(filename string) [][]int {
	input := [][]int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			re := regexp.MustCompile(`\s+`)
			split := re.Split(line, -1)
			var row = []int{}

			for _, v := range split {
				i, _ := strconv.Atoi(v)
				row = append(row, i)
			}
			input = append(input, row)
		}
	}
	return input
}
