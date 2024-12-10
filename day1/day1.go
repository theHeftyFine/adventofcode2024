package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type NumPair struct {
	left  []int
	right []int
}

func Day1(filename string) {
	fmt.Println("Day 1:")
	input1 := Input(filename)
	fmt.Println("Part 1:", part1(input1))
	fmt.Println("Part 2:", part2(input1))
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

func part1(input NumPair) int {
	var sum int = 0
	for i, v := range input.left {
		if len(input.right) > i {
			val1 := float64(v)
			val2 := float64(input.right[i])
			max := math.Max(val1, val2)
			min := math.Min(val1, val2)
			sum = sum + int(max-min)
		}
	}
	return sum
}

func part2(input NumPair) int {
	var sum = 0
	for _, x := range input.left {
		var total = 0
		for _, y := range input.right {
			if x == y {
				total = total + 1
			}
		}
		sum = sum + (total * x)
	}
	return sum
}

func Input(filename string) NumPair {
	input := new(NumPair)
	input.left = []int{}
	input.right = []int{}

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
			val1, _ := strconv.Atoi(split[0])
			val2, _ := strconv.Atoi(split[1])
			input.left = append(input.left, val1)
			input.right = append(input.right, val2)
		}
	}
	slices.Sort(input.left)
	slices.Sort(input.right)
	return *input
}
