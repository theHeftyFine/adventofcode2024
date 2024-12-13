package daydisplay

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
)

type DayRunner[T any] interface {
	Input(filename string) T
	Part1(T, *fyne.Container) int
	Part2(T, *fyne.Container) int
}

type DayDisplay interface {
	Display(filename string) *fyne.Container
}

type BasicDisplay[T any] struct {
	DayRunner[T]
}

func (d *BasicDisplay[T]) exists(item fyne.CanvasObject, items []fyne.CanvasObject) bool {
	for _, it := range items {
		if it == item {
			return true
		}
	}
	return false
}

func (d *BasicDisplay[T]) Widget(filename string) *fyne.Container {
	resultLabel := widget.NewLabel("Result:")
	resultValue := widget.NewLabel("")
	copyButton := widget.NewButton("copy", func() {
		clipboard.Write(clipboard.FmtText, []byte(resultValue.Text))
	})
	resultBar := container.NewHBox(resultLabel, resultValue)
	var cont *fyne.Container
	input := d.Input(filename)

	resultContainer := container.NewCenter()
	button1 := widget.NewButton("Part 1", func() {
		resultValue.SetText(strconv.Itoa(d.Part1(input, resultContainer)))
		if !d.exists(copyButton, resultBar.Objects) {
			resultBar.Add(copyButton)
		}
	})

	button2 := widget.NewButton("Part 2", func() {
		resultValue.SetText(strconv.Itoa(d.Part2(input, resultContainer)))
		if !d.exists(copyButton, resultBar.Objects) {
			resultBar.Add(copyButton)
		}
	})

	buttonRow := container.NewHBox(button1, button2)
	cont = container.NewVBox(buttonRow, resultBar, resultContainer)
	return cont
}
