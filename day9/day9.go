package day9

import (
	"github.com/theheftyfine/adventofcode2024/model"
	"log"
	"os"
	"slices"
	"strconv"

	"fyne.io/fyne/v2"
)

type Block struct {
	size int
	id   int
}

type day struct {
	input []Block
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
				} else {
					b.size -= block.size
					newBlock := block
					newSection := []Block{newBlock, b}
					block.size = 0
					container = slices.Concat(container[:i], newSection, container[i+1:])
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

func input(filename string) []Block {
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
