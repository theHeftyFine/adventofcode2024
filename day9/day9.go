package day9

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Block struct {
	size int
	id   int
}

type fragGridLayout struct {
	Cols int
}

func (g *fragGridLayout) countRows(objects []fyne.CanvasObject) int {
	if g.Cols < 1 {
		g.Cols = 1
	}
	count := 0
	for _, child := range objects {
		if child.Visible() {
			count++
		}
	}

	return int(math.Ceil(float64(count) / float64(g.Cols)))
}

func (g *fragGridLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	rows := g.countRows(objects)
	cols := g.Cols

	cellWidth := float64(size.Width) / float64(rows)
	cellHeight := float64(size.Height) / float64(cols)

	row, col := 0, 0
	i := 0
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		x1 := float32(cellWidth * float64(col))
		y1 := float32(cellHeight * float64(row))
		x2 := float32(cellWidth * float64(col+1))
		y2 := float32(cellHeight * float64(row+1))

		child.Move(fyne.NewPos(x1, y1))
		child.Resize(fyne.NewSize(x2-x1, y2-y1))
		if (i+1)%g.Cols == 0 {
			col++
			row = 0
		} else {
			row++
		}
		i++
	}
}

func (g *fragGridLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	rows := g.countRows(objects)
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}
		minSize = minSize.Max(child.MinSize())
	}

	return fyne.NewSize(minSize.Width*float32(rows), minSize.Height*float32(g.Cols))
}

var colors = []color.RGBA{
	{238, 130, 238, 255}, //violet
	{75, 0, 130, 255},    //indigo
	{0, 0, 255, 255},     //blue
	{0, 128, 0, 255},     //green
	{255, 255, 0, 255},   //yellow
	{255, 165, 130, 255}, //orange
	{255, 0, 0, 255},     //red
}

var gray = color.RGBA{64, 64, 64, 64}

func Day9(filename string) {
	fmt.Println("Day 9:")

	input := Input(filename)

	start1 := time.Now()
	result1 := part1(input)
	elapsed1 := time.Since(start1)
	fmt.Println("Part 1:", result1, "took:", elapsed1.Milliseconds(), "ms")

	start2 := time.Now()
	result2 := part2(input)
	elapsed2 := time.Since(start2)
	fmt.Println("Part 2:", result2, "took:", elapsed2.Milliseconds(), "ms")
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

	var cont *fyne.Container

	buttonGrid := widget.NewButton("Draw grid", func() {
		gridContainer := container.NewCenter()

		rects := createRectangles(input)
		cols := int(math.Floor(math.Sqrt(float64(len(rects)))))
		grid := container.New(&fragGridLayout{cols}, rects...)

		gridContainer.Add(grid)
		cont.Add(gridContainer)
	})

	buttonRow := container.NewHBox(button1, button2, buttonGrid)
	cont = container.NewVBox(buttonRow, resultLabel)
	return cont
}

func createRectangles(input []Block) []fyne.CanvasObject {
	rects := []fyne.CanvasObject{}

	for _, block := range input {
		cl := gray
		if block.id == 0 {
			cl = colors[0]
		} else if block.id > 0 {
			cl = colors[block.id%len(colors)]
		}
		for i := 0; i < block.size; i++ {
			rect := canvas.NewRectangle(cl)
			rect.SetMinSize(fyne.NewSize(2, 2))
			rects = append(rects, rect)
		}
	}
	return rects
}

func part1(input []Block) int {
	out := 0
	container := slices.Clone(input)

	for !checkDefrag(container) {
		var block Block
		for i := len(container) - 1; i > 0; i-- {
			block = container[i]
			if block.id >= 0 {
				head := container[:i]
				tail := container[i+1:]
				container = slices.Concat(head, tail)
				break
			}
		}

		for block.size > 0 {
			for i, b := range container {
				if b.id == -1 {
					if b.size <= block.size {
						block.size -= b.size
						container[i].id = block.id
					} else {
						b.size -= block.size
						newBlock := block
						newSection := []Block{newBlock, b}
						block.size = 0
						container = slices.Concat(container[:i], newSection, container[i+1:])
					}
					break
				}
			}
		}
	}

	pos := 0

	for _, block := range container {
		for i := 0; i < block.size; i++ {
			if block.id != -1 {
				out += pos * block.id
			}
			pos++
		}
	}

	return out
}

func part2(input []Block) int {
	out := 0

	con, moved := moveBlock(input)

	for moved {
		con, moved = moveBlock(con)
	}

	pos := 0

	for _, block := range con {
		for i := 0; i < block.size; i++ {
			if block.id != -1 {
				out += pos * block.id
			}
			pos++
		}
	}

	return out
}

func moveBlock(input []Block) ([]Block, bool) {
	for i, block := range input {
		if block.id < 0 {
			for j := len(input) - 1; j > 0; j-- {
				fit := input[j]
				if fit.id >= 0 && fit.size <= block.size {
					newSection := []Block{fit}
					if fit.size != block.size {
						block.size -= fit.size
						newSection = append(newSection, block)
					}
					// never move blocks to the right
					if j > i {
						empty := Block{id: -1, size: fit.size}
						next := slices.Concat(input[:i], newSection, input[i+1:j], []Block{empty}, input[j+1:])
						return next, true
					}
				}
			}
		}
	}
	return input, false
}

func Input(filename string) []Block {
	out := []Block{}
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	id := 0

	for i, rune := range content {
		s, _ := strconv.Atoi(string(rune))

		block := Block{
			id:   id,
			size: s,
		}

		if i == 0 || i%2 == 0 {
			id++
		} else {
			block.id = -1
		}
		out = append(out, block)
	}
	return out
}

func checkDefrag(blocks []Block) bool {
	end := false
	for _, block := range blocks {
		if end && block.id > -1 {
			return false
		} else if block.id < 0 {
			end = true
		}
	}
	return true
}
