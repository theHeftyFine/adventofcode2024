package model

type Coord struct {
	Y int
	X int
}

func (c Coord) Add(b Coord) Coord {
	return Coord{c.Y + b.Y, c.X + b.X}
}
