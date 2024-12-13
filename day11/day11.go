package day11

import (
	"fyne.io/fyne/v2"
	"github.com/theheftyfine/adventofcode2024/model"
	"log"
	"os"
	"strconv"
	"strings"
)

type day struct {
	input []int
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

func input(filename string) []int {
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

func part1(input []int) int {
	return part(input, 25)
}

func part2(input []int) int {
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
