package day7

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

type Test struct {
	sum   int
	terms []int
}

var ops1 = []func(int, int) int{sum, product}
var ops2 = []func(int, int) int{sum, product, concat}

func Day7(filename string) {
	fmt.Println("Day 7")
	input7 := Input(filename)
	fmt.Println("Part 1:", part1(input7))
	fmt.Println("Part 2:", part2(input7))
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

func part1(input []Test) int {
	return part(input, ops1)
}

func part2(input []Test) int {
	return part(input, ops2)
}

func part(input []Test, ops []func(int, int) int) int {
	out := 0
	for _, p := range input {
		if check(p.sum, p.terms, ops) {
			out += p.sum
		}
	}

	return out
}

func Input(filename string) []Test {
	out := []Test{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			parts := strings.Split(line, ": ")
			sum, _ := strconv.Atoi(parts[0])
			tails := strings.Split(parts[1], " ")

			terms := []int{}
			for _, v := range tails {
				i, _ := strconv.Atoi(v)
				terms = append(terms, i)
			}
			test := Test{
				sum:   sum,
				terms: terms,
			}
			out = append(out, test)
		}
	}
	return out
}

func check(target int, terms []int, ops []func(int, int) int) bool {
	if len(terms) > 1 {
		return checkRecursive(target, terms[0], terms[1:], ops)
	}
	return false
}

func checkRecursive(target int, total int, rem []int, ops []func(int, int) int) bool {
	if len(rem) == 0 {
		return total == target
	} else {
		for _, op := range ops {
			if checkRecursive(target, op(total, rem[0]), rem[1:], ops) {
				return true
			}
		}
		return false
	}
}

func sum(a int, b int) int {
	return a + b
}

func product(a int, b int) int {
	return a * b
}

func concat(a int, b int) int {
	conc := strconv.Itoa(a) + strconv.Itoa(b)
	out, _ := strconv.Atoi(conc)
	return out
}
