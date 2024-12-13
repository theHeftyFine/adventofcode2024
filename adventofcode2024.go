package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/theheftyfine/adventofcode2024/day1"
	"github.com/theheftyfine/adventofcode2024/day10"
	"github.com/theheftyfine/adventofcode2024/day11"
	"github.com/theheftyfine/adventofcode2024/day12"
	"github.com/theheftyfine/adventofcode2024/day2"
	"github.com/theheftyfine/adventofcode2024/day3"
	"github.com/theheftyfine/adventofcode2024/day4"
	"github.com/theheftyfine/adventofcode2024/day5"
	"github.com/theheftyfine/adventofcode2024/day6"
	"github.com/theheftyfine/adventofcode2024/day7"
	"github.com/theheftyfine/adventofcode2024/day8"
	"github.com/theheftyfine/adventofcode2024/day9"
	"golang.design/x/clipboard"
)

var screens = []*fyne.Container{
	day1.Display("input/input1.txt"),
	day2.Display("input/input2.txt"),
	day3.Display("input/input3.txt"),
	day4.Display("input/input4.txt"),
	day5.Display("input/input5.txt"),
	day6.Display("input/input6.txt"),
	day7.Display("input/input7.txt"),
	day8.Display("input/input8.txt"),
	day9.Display("input/input9.txt"),
	day10.Display("input/input10.txt"),
	day11.Display("input/input11.txt"),
	day12.Display("input/input12.txt"),
}

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Panic(err)
	}

	a := app.New()
	w := a.NewWindow("Advent of Code 2024")
	w.Resize(fyne.NewSize(800, 600))
	content := container.NewAppTabs()
	for i, s := range screens {
		content.Append(container.NewTabItem("Day "+strconv.Itoa(i+1), s))
	}
	rows := container.NewVBox(content)
	w.SetContent(rows)

	w.ShowAndRun()

}
