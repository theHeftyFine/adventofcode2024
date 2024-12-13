package day11

import (
	"log"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	daydisplay "github.com/theheftyfine/adventofcode2024/display"
)

type day struct{}

var display = daydisplay.BasicDisplay[[]int]{
	DayRunner: day{},
}

func Display(filename string) *fyne.Container {
	return display.Widget(filename)
}

func (day) Input(filename string) []int {
	out := []int{}
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	digits := strings.Split(string(file), " ")
	for _, d := range digits {
		trim := strings.Trim(d, "\n)")
		i, err := strconv.Atoi(trim)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, i)
	}
	return out
}

func (day) Part1(input []int, cont *fyne.Container) int {
	return part(input, 25)
}

func (day) Part2(input []int, cont *fyne.Container) int {
	return part(input, 25)
}

func part(input []int, times int) int {
	pebblemap := make(map[int]int)
	for _, d := range input {
		pebblemap[d]++
	}
	for i := 0; i < times; i++ {
		nMap := make(map[int]int)
		for d, t := range pebblemap {
			if d == 0 {
				nMap[1] += t
			} else {
				dRunes := []rune(strconv.Itoa(d))
				length := len(dRunes)
				if length%2 == 0 {
					mid := length / 2
					s1 := dRunes[:mid]
					s2 := dRunes[mid:]
					p1, _ := strconv.Atoi(string(s1))
					p2, _ := strconv.Atoi(string(s2))
					nMap[p1] += t
					nMap[p2] += t
				} else {
					nMap[d*2024] += t
				}
			}
		}
		pebblemap = nMap
	}

	out := 0
	count := 0
	for _, t := range pebblemap {
		count++
		out += t
	}
	return out
}

func exists(item fyne.CanvasObject, items []fyne.CanvasObject) bool {
	for _, it := range items {
		if it == item {
			return true
		}
	}
	return false
}
