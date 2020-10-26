package main

type Cell struct {
	Alive      bool
	Coordinate *Coordinate
	Neighbours []*Cell
}

func NewCell(alive bool, coordinate *Coordinate) (cell *Cell) {
	return &Cell{Alive: alive, Coordinate: coordinate}
}

func (cell Cell) String() string {
	if cell.Alive {
		return "x"
	} else {
		return " "
	}
}
