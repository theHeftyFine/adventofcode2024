package day11

import (
	"log"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
)

func Input(filename string) []int {
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

func Widget(filename string) *fyne.Container {
	resultLabel := widget.NewLabel("Result:")
	resultValue := widget.NewLabel("")
	copyButton := widget.NewButton("copy", func() {
		clipboard.Write(clipboard.FmtText, []byte(resultValue.Text))
	})
	resultBar := container.NewHBox(resultLabel, resultValue)
	var cont *fyne.Container
	input := Input(filename)
	button1 := widget.NewButton("Part 1", func() {
		resultValue.SetText(strconv.Itoa(part(input, 25)))
		if !exists(copyButton, resultBar.Objects) {
			resultBar.Add(copyButton)
		}
	})

	button2 := widget.NewButton("Part 2", func() {
		resultValue.SetText(strconv.Itoa(part(input, 75)))
		if !exists(copyButton, resultBar.Objects) {
			resultBar.Add(copyButton)
		}
	})

	buttonRow := container.NewHBox(button1, button2)
	cont = container.NewVBox(buttonRow, resultBar)
	return cont
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
