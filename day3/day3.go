package day3

import (
	"github.com/theheftyfine/adventofcode2024/model"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type day struct {
	input string
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

func part1(input string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	results := re.FindAllString(input, -1)

	var sum = 0

	for _, v := range results {
		sum += calcMul(v)
	}
	return sum
}

func part2(input string) int {
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

func input(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
