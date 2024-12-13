package day9

import (
	"image/color"
	"log"
	"math"
	"os"
	"slices"
	"strconv"

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

		child.Move(fyne.NewPos(y1, x1))
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

func Display(filename string) *fyne.Container {
	resultLabel := widget.NewLabel("")
	input := Input(filename)
	button1 := widget.NewButton("Part 1", func() {
		resultLabel.SetText("Result: " + strconv.Itoa(part1(input)))
	})

	button2 := widget.NewButton("Part 2", func() {
		resultLabel.SetText("Result: " + strconv.Itoa(part2(input)))
	})

	var cont *fyne.Container
	gridContainer := container.NewCenter()

	button1draw := widget.NewButton("Simulate Part 1 ", func() {
		rectGrid := displayGrid(input, gridContainer)
		resultLabel.SetText("Result: " + strconv.Itoa(part1draw(input, rectGrid)))
	})

	buttonRow := container.NewHBox(button1, button2, button1draw)
	cont = container.NewVBox(buttonRow, resultLabel, gridContainer)
	return cont
}

func displayGrid(input []Block, cont *fyne.Container) *fyne.Container {
	cont.RemoveAll()

	rects := createRectangles(input)
	cols := int(math.Floor(math.Sqrt(float64(len(rects)))))
	grid := container.New(&fragGridLayout{cols}, rects...)

	cont.Add(grid)
	return grid
}

func updateGrid(cont *fyne.Container, pos int, blocks ...Block) {
	if cont != nil {
		rects := cont.Objects
		index := pos
		for _, block := range blocks {
			clr := gray
			if block.id >= 0 {
				clr = colors[block.id%len(colors)]
				for i := 0; i < block.size; i++ {
					rect, ok := rects[index].(*canvas.Rectangle)
					if ok && rect.FillColor != clr {
						rect.FillColor = clr
					}
					index++
				}
			}
		}
		cont.Refresh()
	}
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
		container = moveBlocks(container, nil)
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

func part1draw(input []Block, cont *fyne.Container) int {
	out := 0
	container := slices.Clone(input)

	for !checkDefrag(container) {
		container = moveBlocks(container, cont)
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

func moveBlocks(container []Block, cont *fyne.Container) []Block {
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
		pos := 0
		for i, b := range container {
			if b.id == -1 {
				if b.size <= block.size {
					block.size -= b.size
					container[i].id = block.id
					updateGrid(cont, pos, container[i])
				} else {
					b.size -= block.size
					newBlock := block
					newSection := []Block{newBlock, b}
					block.size = 0
					container = slices.Concat(container[:i], newSection, container[i+1:])
					updateGrid(cont, pos, newBlock, b)
				}
				break
			}
			pos += b.size
		}
	}
	return container
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
