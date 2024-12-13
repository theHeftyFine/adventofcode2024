package day1

import (
	"bufio"
	"github.com/theheftyfine/adventofcode2024/model"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type day struct {
	input NumPair
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

type NumPair struct {
	left  []int
	right []int
}

func part1(input NumPair) int {
	sum := 0
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
	sum := 0
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

func input(filename string) NumPair {
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
