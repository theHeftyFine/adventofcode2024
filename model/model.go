package model

type Coord struct {
	Y int
	X int
}

type Tile struct {
	Loc  Coord
	Crop rune
	Set  bool
}

func (c Coord) New(y int, x int) Coord {
	return Coord{y, x}
}

func (c Coord) Add(b Coord) Coord {
	return Coord{c.Y + b.Y, c.X + b.X}
}

func (c Coord) Flatten(maxH int) int {
	return (c.Y * maxH) + c.X
}

func (c Coord) Includes(Coords []Coord) bool {
	for _, co := range Coords {
		if c == co {
			return true
		}
	}
	return false
}

func (c Coord) Turn() Coord {
	return Coord{Y: c.X, X: c.Y}
}

func (c Coord) Negate() Coord {
	return Coord{Y: c.Y * -1, X: c.X * -1}
}

func (c Coord) Border(dir Coord, fields map[Coord]Tile) bool {
	return fields[c.Add(dir)].Set
}
