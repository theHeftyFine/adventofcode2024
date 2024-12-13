package day12

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"
	"os"

	"fyne.io/fyne/v2"
	daydisplay "github.com/theheftyfine/adventofcode2024/display"
	"github.com/theheftyfine/adventofcode2024/model"
)

type tile = model.Tile

type Field struct {
	tiles map[coord]tile
	crop  rune
}

type coord = model.Coord

// func (c coord) Add(b coord) coord {
// 	return coord{c.y + b.y, c.x + b.x}
// }

// func (c coord) Flatten(maxH int) int {
// 	return (c.y * maxH) + c.x
// }

// func (c coord) Includes(coords []coord) bool {
// 	for _, co := range coords {
// 		if c == co {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (c coord) Turn() coord {
// 	return coord{y: c.x, x: c.y}
// }

// func (c coord) Negate() coord {
// 	return coord{y: c.y * -1, x: c.x * -1}
// }

// func (c coord) Border(dir coord, fields map[coord]tile) bool {
// 	return fields[c.Add(dir)].set
// }

var dirs = []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type day struct{}

var display = daydisplay.BasicDisplay[[][]tile]{
	DayRunner: day{},
}

func Display(filename string) *fyne.Container {
	return display.Widget(filename)
}

func (day) Input(filename string) [][]tile {
	out := [][]tile{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	y := 0

	for scanner.Scan() {
		row := []tile{}
		for x, p := range scanner.Text() {
			row = append(row, tile{coord{y, x}, p, true})
		}
		out = append(out, row)
		y++
	}
	return out
}

func (day) Part1(input [][]tile, cont *fyne.Container) int {
	return calc(input, cont, calcPerimiter)
}

func (day) Part2(input [][]tile, cont *fyne.Container) int {
	return calc(input, cont, calcPerimiter2)
}

func calc(input [][]tile, cont *fyne.Container, f func(field map[coord]tile) (int, []coord)) int {
	out := 0
	covered := []tile{}
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
		for _, t := range row {
			if includes(t, covered) {
				continue
			}
			field := map[coord]tile{} //[]tile{tile}
			field[t.Loc] = t
			perimiter := []tile{t}
			crop := t.Crop

			for len(perimiter) > 0 {
				nPerimiter := []tile{}
				for _, ti := range perimiter {
					for _, dir := range dirs {
						nPos := dir.Add(ti.Loc)
						if inBound(nPos, input) {
							ntile := input[nPos.Y][nPos.X]
							if ntile.Crop == crop && !includesMap(ntile, field) {
								nPerimiter = append(nPerimiter, ntile)
								field[ntile.Loc] = ntile
								covered = append(covered, ntile)
							}
						}
					}
				}
				perimiter = nPerimiter
			}
			fields = append(fields, Field{field, crop})
		}
	}

	for _, field := range fields {
		surface := len(field.tiles)
		perimiter, _ := f(field.tiles)
		cost := surface * perimiter
		fmt.Println("- A region of", string(field.crop), "plants with price", surface, "*", perimiter, "=", cost)
		out += cost
	}

	return out
}

func calcPerimiter(field map[coord]tile) (int, []coord) {
	count := 0
	border := []coord{}
	for _, t := range field {
		for _, dir := range dirs {
			locCount := 0
			bord := field[dir.Add(t.Loc)]
			if !bord.Set {
				locCount++
			}
			if locCount > 0 {
				border = append(border, t.Loc)
			}
			count += locCount
		}
	}
	return count, border
}

func calcPerimiter2(field map[coord]tile) (int, []coord) {
	count := 0
	border := []coord{}
	dirChecked := map[coord][]coord{}
	// check every field to see if it has a border
	for _, t := range field {
		// check in all four directions of the field
		for _, dir := range dirs {
			// only count a field if it has not already been counted in that direction
			// (if we get to a tile that has been checked, it is included in a bulk)
			if !t.Loc.Includes(dirChecked[dir]) {
				locCount := 0

				// the field in that direction is not set, so it is a border
				if !t.Loc.Border(dir, field) {
					locCount++

					// check in both directions, parallel to the border
					checkDirs := []coord{dir.Turn(), dir.Turn().Negate()}
					for _, cDir := range checkDirs {
						// start with the first tile (this will duplicate it in the checked tiles, but it doesnt matter)
						ctile := t
						// as long as the next tile is checked, and a border
						for ctile.Set && !ctile.Loc.Border(dir, field) {
							// add this tile to the checked tiles in the currect direction
							dirChecked[dir] = append(dirChecked[dir], ctile.Loc)
							// set the next tile to be the next tile along the edge
							ctile = field[ctile.Loc.Add(cDir)]
						}
					}
				}
				if locCount > 0 {
					border = append(border, t.Loc)
				}
				count += locCount
			}
		}
	}
	return count, border
}

func includes(tile tile, tiles []tile) bool {
	for _, t := range tiles {
		if t == tile {
			return true
		}
	}
	return false
}

func includesMap(tile tile, tiles map[coord]tile) bool {
	for _, t := range tiles {
		if t == tile {
			return true
		}
	}
	return false
}

func inBound(t coord, input [][]tile) bool {
	return t.Y >= 0 && t.Y < len(input) && t.X >= 0 && t.X < len(input[0])
}

// func findField(y int, x int, tiles []tile) bool {
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

func getNeighbors(field map[coord]tile, c coord, dir coord) []tile {
	n1 := field[c.Add(dir.Turn())]
	n2 := field[c.Add(dir.Turn().Negate())]
	neighbours := []tile{}
	if n1.Set {
		neighbours = append(neighbours, n1)
	}
	if n2.Set {
		neighbours = append(neighbours, n2)
	}
	return neighbours
}
