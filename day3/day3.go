package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Day3(filename string) {
	fmt.Println("Day 3:")
	input3 := Input(filename)
	fmt.Println("when adding up all multiplication, as indicated by the method mul(x, y), the answer", part1(input3))
	fmt.Println("However, when the method don't() stops processing, and the method do() starts it, the answer is", part2(input3))
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

func Input(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
