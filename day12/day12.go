package day12

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	daydisplay "github.com/theheftyfine/adventofcode2024/display"
)

type Tile struct {
	loc  coord
	crop rune
	set  bool
}

type Field struct {
	tiles map[coord]Tile
	crop  rune
}

type coord struct {
	y int
	x int
}

func (c coord) Add(b coord) coord {
	return coord{c.y + b.y, c.x + b.x}
}

func (c coord) Flatten(maxH int) int {
	return (c.y * maxH) + c.x
}

func (c coord) Includes(coords []coord) bool {
	for _, co := range coords {
		if c == co {
			return true
		}
	}
	return false
}

func (c coord) Turn() coord {
	return coord{y: c.x, x: c.y}
}

func (c coord) Negate() coord {
	return coord{y: c.y * -1, x: c.x * -1}
}

var dirs = []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type day struct{}

var display = daydisplay.BasicDisplay[[][]Tile]{
	DayRunner: day{},
}

func Display(filename string) *fyne.Container {
	return display.Widget(filename)
}

func (day) Input(filename string) [][]Tile {
	out := [][]Tile{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	y := 0

	for scanner.Scan() {
		row := []Tile{}
		for x, p := range scanner.Text() {
			row = append(row, Tile{coord{y, x}, p, true})
		}
		out = append(out, row)
		y++
	}
	return out
}

func (day) Part1(input [][]Tile, cont *fyne.Container) int {
	return calc(input, cont, calcPerimiter)
}

func (day) Part2(input [][]Tile, cont *fyne.Container) int {
	return calc(input, cont, calcPerimiter2)
}

func calc(input [][]Tile, cont *fyne.Container, f func(field map[coord]Tile) (int, []coord)) int {
	out := 0
	covered := []Tile{}
	fields := []Field{}

	// plots := []fyne.CanvasObject{}

	// cont.RemoveAll()

	// for _, row := range input {
	// 	for i := 0; i < len(row); i++ {
	// 		plots = append(plots, canvas.NewText(".", color.White))
	// 	}
	// }

	// grid := container.NewGridWithColumns(len(input), plots...)
	// cont.Add(grid)

	for _, row := range input {
		for _, tile := range row {
			if includes(tile, covered) {
				continue
			}
			field := map[coord]Tile{} //[]Tile{tile}
			field[tile.loc] = tile
			perimiter := []Tile{tile}
			crop := tile.crop

			for len(perimiter) > 0 {
				nPerimiter := []Tile{}
				for _, tile := range perimiter {
					for _, dir := range dirs {
						nPos := dir.Add(tile.loc)
						if inBound(nPos, input) {
							nTile := input[nPos.y][nPos.x]
							if nTile.crop == crop && !includesMap(nTile, field) {
								nPerimiter = append(nPerimiter, nTile)
								field[nTile.loc] = nTile
								covered = append(covered, nTile)
							}
						}
					}
				}
				perimiter = nPerimiter
			}
			fields = append(fields, Field{field, crop})
		}
	}

	progress := widget.NewProgressBar()
	progCont := container.NewHBox(progress)
	cont.Add(progCont)

	for i, field := range fields {
		surface := len(field.tiles)
		perimiter, _ := f(field.tiles)
		cost := surface * perimiter
		fmt.Println("- A region of", string(field.crop), "plants with price", surface, "*", perimiter, "=", cost)
		progress.SetValue(float64(i+1) / float64(len(fields)))
		out += cost
	}
	cont.Remove(progress)

	return out
}

func calcPerimiter(field map[coord]Tile) (int, []coord) {
	count := 0
	border := []coord{}
	for _, tile := range field {
		for _, dir := range dirs {
			locCount := 0
			bord := field[dir.Add(tile.loc)]
			if !bord.set {
				locCount++
			}
			if locCount > 0 {
				border = append(border, tile.loc)
			}
			count += locCount
		}
	}
	return count, border
}

func calcPerimiter2(field map[coord]Tile) (int, []coord) {
	count := 0
	border := []coord{}
	dirChecked := map[coord][]coord{}
	for _, tile := range field {
		for _, dir := range dirs {
			locCount := 0
			c := dir.Add(tile.loc)
			bord := field[c]
			if !bord.set && !c.Includes(dirChecked[dir]) {
				locCount++
				checkDirs := []coord{dir.Turn(), dir.Turn().Negate()}
				for _, cDir := range checkDirs {
					cTile := tile
					for cTile.set && !field[dir].set {
						dirChecked[dir] = append(dirChecked[dir], cTile.loc)
						cTile = field[cTile.loc.Add(cDir)]
					}
				}
			}
			if locCount > 0 {
				border = append(border, tile.loc)
			}
			count += locCount
		}
	}
	return count, border
}

func includes(tile Tile, tiles []Tile) bool {
	for _, t := range tiles {
		if t == tile {
			return true
		}
	}
	return false
}

func includesMap(tile Tile, tiles map[coord]Tile) bool {
	for _, t := range tiles {
		if t == tile {
			return true
		}
	}
	return false
}

func inBound(tile coord, input [][]Tile) bool {
	return tile.y >= 0 && tile.y < len(input) && tile.x >= 0 && tile.x < len(input[0])
}

// func findField(y int, x int, tiles []Tile) bool {
// 	for _, tile := range tiles {
// 		if y == tile.y && x == tile.x {
// 			return true
// 		}
// 	}
// 	return false
// }

func randomColor() color.Color {
	r := uint8(rand.Float32() * 255)
	g := uint8(rand.Float32() * 255)
	b := uint8(rand.Float32() * 255)
	return color.RGBA{r, g, b, 255}
}

func getNeighbors(field map[coord]Tile, c coord, dir coord) []Tile {
	n1 := field[c.Add(dir.Turn())]
	n2 := field[c.Add(dir.Turn().Negate())]
	neighbours := []Tile{}
	if n1.set {
		neighbours = append(neighbours, n1)
	}
	if n2.set {
		neighbours = append(neighbours, n2)
	}
	return neighbours
}
