package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/theheftyfine/adventofcode2024/model"
)

type tile = model.Tile

type field struct {
	tiles map[coord]tile
	crop  rune
}

type coord = model.Coord

var dirs = []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type day struct {
	input [][]tile
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

func input(filename string) [][]tile {
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

func part1(input [][]tile) int {
	return calc(input, calcPerimiter)
}

func part2(input [][]tile) int {
	return calc(input, calcPerimiter2)
}

func calc(input [][]tile, f func(field map[coord]tile) (int, []coord)) int {
	out := 0
	covered := []tile{}
	fields := []field{}

	for _, row := range input {
		for _, t := range row {
			if includes(t, covered) {
				continue
			}
			f := map[coord]tile{} //[]tile{tile}
			f[t.Loc] = t
			perimiter := []tile{t}
			crop := t.Crop

			for len(perimiter) > 0 {
				nPerimiter := []tile{}
				for _, ti := range perimiter {
					for _, dir := range dirs {
						nPos := dir.Add(ti.Loc)
						if inBound(nPos, input) {
							ntile := input[nPos.Y][nPos.X]
							if ntile.Crop == crop && !includesMap(ntile, f) {
								nPerimiter = append(nPerimiter, ntile)
								f[ntile.Loc] = ntile
								covered = append(covered, ntile)
							}
						}
					}
				}
				perimiter = nPerimiter
			}
			fields = append(fields, field{f, crop})
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
