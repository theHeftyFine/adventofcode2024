package day3

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	daydisplay "github.com/theheftyfine/adventofcode2024/display"
)

type day struct{}

var display = daydisplay.BasicDisplay[string]{
	DayRunner: day{},
}

func Display(filename string) *fyne.Container {
	return display.Widget(filename)
}

func (day) Part1(input string, cont *fyne.Container) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	results := re.FindAllString(input, -1)

	var sum = 0

	for _, v := range results {
		sum += calcMul(v)
	}
	return sum
}

func (day) Part2(input string, cont *fyne.Container) int {
	mul := `mul\(\d{1,3},\d{1,3}\)`
	do := `do\(\)`
	dont := `don\'t\(\)`

	mulRe := regexp.MustCompile(mul)
	doRe := regexp.MustCompile(do)
	dontRe := regexp.MustCompile(dont)

	re := regexp.MustCompile(mul + "|" + do + "|" + dont)

	var consume = strings.Clone(input)

	var index = re.FindStringIndex(consume)

	var parse bool = true

	var sum = 0

	for index != nil {
		match := re.FindString(consume)
		if mulRe.MatchString(match) && parse {
			sum += calcMul(match)
		} else if doRe.MatchString(match) && !parse {
			parse = true
		} else if dontRe.MatchString(match) && parse {
			parse = false
		}

		i := index[1]
		consume = consume[i:]
		index = re.FindStringIndex(consume)
	}

	return sum
}

func calcMul(mul string) int {
	reDigit := regexp.MustCompile(`\d{1,3}`)
	digits := reDigit.FindAllString(mul, -1)
	var product = 1
	for _, d := range digits {
		i, _ := strconv.Atoi(d)
		product = product * i
	}
	return product
}

func (day) Input(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
