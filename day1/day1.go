package day1

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"

	"fyne.io/fyne/v2"
	daydisplay "github.com/theheftyfine/adventofcode2024/display"
)

type NumPair struct {
	left  []int
	right []int
}

type day struct{}

var display = daydisplay.BasicDisplay[NumPair]{
	DayRunner: day{},
}

func Display(filename string) *fyne.Container {
	return display.Widget(filename)
}

func (day) Part1(input NumPair, cont *fyne.Container) int {
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

func (day) Part2(input NumPair, cont *fyne.Container) int {
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

func (day) Input(filename string) NumPair {
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
