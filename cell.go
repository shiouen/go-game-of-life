package main

import "fmt"

type Cell struct {
	Alive      bool
	Coordinate *Coordinate
	Index      int
	Neighbours []*Cell
}

func NewCell(alive bool, coordinate *Coordinate, index int) (cell *Cell) {
	return &Cell{Alive: alive, Coordinate: coordinate, Index: index}
}

func (cell Cell) String() string {
	//if cell.Alive {
	//	return "x"
	//} else {
	//	return " "
	//}
	return fmt.Sprintf("(%d, %d)", cell.Coordinate.X, cell.Coordinate.Y)
}
