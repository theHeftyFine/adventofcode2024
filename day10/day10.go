package day10

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Node struct {
	x    int
	y    int
	h    int
	prev *Node
	next []Node
}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func Day10(filename string) {
	fmt.Println("Day 10")
	input := Input(filename)
	part1, part2, _ := part(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
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

	mapContainer := container.NewCenter()

	buttonDraw := widget.NewButton("Draw map", func() {
		_, _, routes := part(input)
		rects := drawRectangles(input)
		mapContainer.Add(fillWithRects(rects, len(input)))
		for _, route := range routes {
			go func() {
				simulateRoute(rects, route, input)
			}()
		}
	})

	buttonRow := container.NewHBox(button1, button2, buttonDraw)

	rows := container.NewVBox(buttonRow, resultLabel, mapContainer)
	return rows
}

func simulateRoute(rects []*canvas.Rectangle, route []Node, input [][]int) {
	resetRectangles(rects, input)
	for _, r := range route {
		col := uint8((255/10)*r.h + 1)
		h := color.RGBA{col, 0, 0, 255}
		if r.h == 9 {
			h.G = 255
		}
		rect := canvas.NewRectangle(h)
		rect.SetMinSize(fyne.NewSize(10, 10))

		index := (r.y * (len(input))) + r.x
		rects[index].FillColor = h
		rects[index].Refresh()
		time.Sleep(50 * time.Millisecond)
	}

}

func fillWithRects(rects []*canvas.Rectangle, cols int) *fyne.Container {
	grid := container.New(layout.NewGridLayout(cols))
	for _, rect := range rects {
		grid.Add(rect)
	}
	return grid
}

func resetRectangles(rects []*canvas.Rectangle, input [][]int) {
	for j, row := range input {
		for i, cell := range row {
			index := (j * (len(input))) + i
			col := uint8((256 / 10) * (cell + 1))
			h := color.RGBA{0, col, 0, 255}
			if rects[index].FillColor != h {
				rects[index].FillColor = h
				rects[index].Refresh()
			}
		}
	}
}

func drawRectangles(input [][]int) []*canvas.Rectangle {
	rects := []*canvas.Rectangle{}

	for _, row := range input {
		for _, cell := range row {
			col := uint8((256 / 10) * (cell + 1))
			h := color.RGBA{0, col, 0, 255}
			rect := canvas.NewRectangle(h)
			rect.SetMinSize(fyne.NewSize(10, 10))
			rects = append(rects, rect)
		}
	}
	return rects
}

func part1(input [][]int) int {
	out, _, _ := part(input)
	return out
}

func part2(input [][]int) int {
	_, out, _ := part(input)
	return out
}

func part(input [][]int) (int, int, [][]Node) {
	out := 0
	heads := []Node{}

	v := len(input)
	w := len(input[0])

	for y, row := range input {
		for x, h := range row {
			if h == 0 {
				node := Node{
					x:    x,
					y:    y,
					h:    h,
					prev: nil,
					next: []Node{},
				}
				heads = append(heads, node)
			}
		}
	}

	moved := true

	for moved {
		moved = false
		nNodes := []Node{}
		for _, node := range heads {
			for _, dir := range dirs {
				ny := node.y + dir[0]
				nx := node.x + dir[1]
				if ny >= 0 && ny < v && nx >= 0 && nx < w {
					h := input[ny][nx]
					if node.h+1 == h {
						if node.prev == nil || !(node.prev.x == nx && node.prev.y == ny) {
							newNode := Node{
								x:    nx,
								y:    ny,
								h:    h,
								prev: &node,
								next: []Node{},
							}
							node.next = append(node.next, newNode)
							if !contains(node, nNodes) {
								nNodes = append(nNodes, newNode)
							}
						}
					}
				}
			}
		}
		if len(nNodes) > 0 {
			moved = true
			heads = nNodes
		}
	}

	trails := map[string][][]int{}
	maps := [][][]rune{}
	routes := [][]Node{}

	for _, head := range heads {
		root, mp, route := traceMap(head, input)
		routes = append(routes, route)

		maps = append(maps, mp)

		xs := strconv.Itoa(root.x)
		ys := strconv.Itoa(root.y)
		coord := []int{head.y, head.x}
		key := ys + "-" + xs
		if !containsCoord(coord, trails[key]) {
			trails[key] = append(trails[key], coord)
		}
	}

	for _, v := range trails {
		out += len(v)
	}

	return out, len(maps), routes
}

func traceMap(head Node, input [][]int) (Node, [][]rune, []Node) {
	mp := [][]rune{}
	routes := []Node{}

	for i := 0; i < len(input); i++ {
		row := []rune{}
		for j := 0; j < len(input[0]); j++ {
			row = append(row, '.')
		}
		mp = append(mp, row)
	}

	mp[head.y][head.x] = rune(head.h) + '0'
	routes = append(routes, head)

	root := head.prev
	for root.h != 0 {
		mp[root.y][root.x] = rune(root.h) + '0'
		routes = append(routes, *root)
		root = root.prev
	}
	mp[root.y][root.x] = rune(root.h) + '0'
	routes = append(routes, *root)

	slices.Reverse(routes)
	return *root, mp, routes
}

func Input(filename string) [][]int {
	out := [][]int{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := []int{}
		for _, r := range scanner.Text() {
			i, _ := strconv.Atoi(string(r))
			row = append(row, i)
		}
		out = append(out, row)
	}
	return out
}

func contains(node Node, nodes []Node) bool {
	for _, n := range nodes {
		if node.x == n.x && node.y == n.y {
			return true
		}
	}
	return false
}

func containsCoord(coord []int, coords [][]int) bool {
	for _, c := range coords {
		if coord[0] == c[0] && coord[1] == c[1] {
			return true
		}
	}
	return false
}
