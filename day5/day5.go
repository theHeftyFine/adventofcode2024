package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Day5(filename string) {
	fmt.Println("Day 5")
	rules5, updates5 := Input(filename)
	fmt.Println("part 1:", part1(rules5, updates5))
	fmt.Println("Part 2:", part2(rules5, updates5))
}

func Widget(filename string) *fyne.Container {
	resultLabel := widget.NewLabel("")
	rules, updates := Input(filename)
	button1 := widget.NewButton("Part 1", func() {
		resultLabel.SetText("Result: " + strconv.Itoa(part1(rules, updates)))
	})

	button2 := widget.NewButton("Part 2", func() {
		resultLabel.SetText("Result: " + strconv.Itoa(part2(rules, updates)))
	})

	buttonRow := container.NewHBox(button1, button2)
	return container.NewVBox(buttonRow, resultLabel)
}

func Input(filename string) (map[int][]int, [][]int) {
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

	return ruleMap, updates

}

func part1(rules map[int][]int, updates [][]int) int {
	var sum = 0
	for _, row := range updates {
		if checkRow(row, rules) {
			sum += getMid(row)
		}
	}
	return sum
}

func part2(rules map[int][]int, updates [][]int) int {
	var sum = 0
	for _, row := range updates {
		if !checkRow(row, rules) {
			sum += getMid(rearrange(row, rules))
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
