package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	daydisplay "github.com/theheftyfine/adventofcode2024/display"
)

type Test struct {
	sum   int
	terms []int
}

var ops1 = []func(int, int) int{sum, product}
var ops2 = []func(int, int) int{sum, product, concat}

type day struct{}

var display = daydisplay.BasicDisplay[[]Test]{
	DayRunner: day{},
}

func Display(filename string) *fyne.Container {
	return display.Widget(filename)
}

func (day) Part1(input []Test, cont *fyne.Container) int {
	return part(input, ops1)
}

func (day) Part2(input []Test, cont *fyne.Container) int {
	return part(input, ops2)
}

func (day) Input(filename string) []Test {
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

func part(input []Test, ops []func(int, int) int) int {
	out := 0
	for _, p := range input {
		if check(p.sum, p.terms, ops) {
			out += p.sum
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
