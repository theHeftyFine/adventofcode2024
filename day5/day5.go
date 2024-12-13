package day5

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	daydisplay "github.com/theheftyfine/adventofcode2024/display"
)

type input struct {
	rules   map[int][]int
	updates [][]int
}

type day struct{}

var display = daydisplay.BasicDisplay[input]{
	DayRunner: day{},
}

func Display(filename string) *fyne.Container {
	return display.Widget(filename)
}

func (day) Input(filename string) input {
	var ruleMap = make(map[int][]int)
	var updates = [][]int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var rules = true

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			rules = false
		} else if rules {
			parts := strings.Split(line, "|")
			key, _ := strconv.Atoi(parts[0])
			val, _ := strconv.Atoi(parts[1])

			ruleMap[key] = append(ruleMap[key], val)
		} else {
			updateLine := strings.Split(line, ",")
			var update = []int{}
			for _, v := range updateLine {
				i, _ := strconv.Atoi(v)
				update = append(update, i)
			}
			updates = append(updates, update)
		}
	}

	return input{ruleMap, updates}

}

func (day) Part1(update input, cont *fyne.Container) int {
	var sum = 0
	for _, row := range update.updates {
		if checkRow(row, update.rules) {
			sum += getMid(row)
		}
	}
	return sum
}

func (day) Part2(update input, cont *fyne.Container) int {
	var sum = 0
	for _, row := range update.updates {
		if !checkRow(row, update.rules) {
			sum += getMid(rearrange(row, update.rules))
		}
	}
	return sum
}

func checkRow(row []int, rules map[int][]int) bool {
	var correct = true
	for i, val := range row {
		if i > 0 {
			before := rules[val]
			preceding := row[:i]
			for _, p := range before {
				if contains(preceding, p) {
					correct = false
				}
			}
		}
	}
	return correct
}

func rearrange(row []int, rules map[int][]int) []int {
	var new = []int{}

	for _, val := range row {
		if len(new) == 0 {
			new = append(new, val)
		} else {
			rule := rules[val]
			var pre = []int{}
			var post = []int{}
			for _, p := range new {
				if contains(rule, p) {
					post = append(post, p)
				} else {
					pre = append(pre, p)
				}
			}
			new = slices.Concat(pre, []int{val}, post)
		}
	}
	return new
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getMid(row []int) int {
	if len(row)%2 == 0 {
		return 0
	}
	return row[int((len(row)-1)/2)]
}
