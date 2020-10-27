package main

import (
	"math"
)

type Grid struct {
	Cells  []*Cell
	Height int
	Width  int
	Size   int
}

func NewGrid(width int, height int) *Grid {
	var size = width * height
	var cells []*Cell

	for i := 0; i < size; i++ {
		var coordinate = &Coordinate{
			X: i % width,
			Y: int(math.Floor(float64(i) / float64(width))),
		}

		cells = append(cells, NewCell(true, coordinate, i))
	}

	var grid = &Grid{Cells: cells, Height: height, Width: width, Size: size}

	for _, cell := range grid.Cells {
		grid.SetNeighbours(cell)
		//fmt.Printf("| index: %d, neightbours: %s |", cell.Index, cell.Neighbours)
	}

	return grid
}

func (grid Grid) SetNeighbours(cell *Cell) {
	var ci = cell.Index
	var above = ci + grid.Width
	var below = ci - grid.Width

	var indexes = []int{
		above - 1, above, above + 1,
		ci - 1, ci + 1,
		below - 1, below, below + 1,
	}

	var cx = cell.Coordinate.X
	var cy = cell.Coordinate.Y

	for _, ni := range indexes {
		if ni >= 0 && ni < grid.Size {
			var neighbour = grid.Cells[ni]
			var nx = neighbour.Coordinate.X
			var ny = neighbour.Coordinate.Y

			if math.Abs(float64(nx-cx)) <= 1 && math.Abs(float64(ny-cy)) <= 1 {
				cell.Neighbours = append(cell.Neighbours, grid.Cells[ni])
			}
		}
	}
}
